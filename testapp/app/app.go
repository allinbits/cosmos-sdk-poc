package app

import (
	"flag"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/fdymylja/tmos/runtime"
	testmodule "github.com/fdymylja/tmos/testapp/module"
	authn2 "github.com/fdymylja/tmos/x/authn"
	bank2 "github.com/fdymylja/tmos/x/bank"
	distribution2 "github.com/fdymylja/tmos/x/distribution"
	"github.com/spf13/viper"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/config"
	tmflags "github.com/tendermint/tendermint/libs/cli/flags"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/proxy"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "testapp/app/config/config.toml", "Path to config.toml")
}

func NewApp() abci.Application {
	rtb := runtime.NewBuilder()
	authentication := authn2.NewModule()
	rtb.AddModule(authentication)
	rtb.SetDecoder(authentication.GetTxDecoder())
	rtb.AddModule(bank2.NewModule())
	rtb.AddModule(distribution2.NewModule())
	rtb.AddModule(testmodule.NewModule())
	rt, err := rtb.Build()
	if err != nil {
		panic(err)
	}
	tmApp := runtime.NewABCIApplication(rt)
	return tmApp
}

func New() {
	rtb := runtime.NewBuilder()
	authentication := authn2.NewModule()
	rtb.AddModule(authentication)
	rtb.SetDecoder(authentication.GetTxDecoder())
	rtb.AddModule(bank2.NewModule())
	rtb.AddModule(distribution2.NewModule())
	rtb.AddModule(testmodule.NewModule())
	rt, err := rtb.Build()
	if err != nil {
		panic(err)
	}
	tmApp := runtime.NewABCIApplication(rt)

	tmNode, err := newTendermint(tmApp, configFile)
	if err != nil {
		panic(err)
	}
	err = tmNode.Start()
	if err != nil {
		panic(err)
	}
	defer tmNode.Stop()
	defer tmNode.Wait()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	os.Exit(0)
}

func newTendermint(app abci.Application, confPath string) (*node.Node, error) {
	conf := config.DefaultConfig()
	conf.RootDir = filepath.Dir(confPath)
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}
	if err := conf.ValidateBasic(); err != nil {
		return nil, err
	}
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stderr))
	logger, err := tmflags.ParseLogLevel(config.DefaultLogLevel, logger, config.DefaultLogLevel)
	if err != nil {
		return nil, err
	}
	pv := privval.LoadFilePV(conf.PrivValidatorKeyFile(), conf.PrivValidatorStateFile())
	nodeKey, err := p2p.LoadNodeKey(conf.NodeKeyFile())
	if err != nil {
		return nil, err
	}
	n, err := node.NewNode(
		conf,
		pv,
		nodeKey,
		proxy.NewLocalClientCreator(app),
		node.DefaultGenesisDocProviderFunc(conf),
		node.DefaultDBProvider,
		node.DefaultMetricsProvider(conf.Instrumentation),
		logger)
	if err != nil {
		return nil, err
	}
	return n, nil
}
