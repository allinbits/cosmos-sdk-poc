package spfcli

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

const (
	DefaultConfigDirectory = ".spfcli_config"
	ConfigName             = "config.yaml"
	ConfigDirectoryEnv     = "SPFCLI_CONFIG_DIR"
)

type Config struct {
	TendermintRPC string
}

func (c *Config) Validate() error {
	if c.TendermintRPC == "" {
		return fmt.Errorf("missing tendermint RPC")
	}
	return nil
}

func GetConfig() (*Config, error) {
	// first check if env is set
	path, ok := os.LookupEnv(ConfigDirectoryEnv)
	if !ok {
		path = DefaultConfigDirectory
	}
	configPath := filepath.Join(path, ConfigName)
	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	// unmarshal configs
	b, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, err
	}
	conf := new(Config)
	err = yaml.Unmarshal(b, conf)
	if err != nil {
		return nil, err
	}
	err = conf.Validate()
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func ConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "handle configurations",
	}
	cmd.AddCommand()
	return cmd
}

func ConfigListCmd() *cobra.Command {
	return &cobra.Command{
		Use: "list",
	}
}

func ConfigUseContext() *cobra.Command {
	return &cobra.Command{
		Use: "use-context",
	}
}
