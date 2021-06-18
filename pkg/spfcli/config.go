package spfcli

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"k8s.io/klog/v2"
)

const (
	DefaultConfigDirectory = ".spfcli_config"
	ConfigName             = "config.yaml"
	ConfigDirectoryEnv     = "SPFCLI_CONFIG_DIR"
)

type Config struct {
	// Name is the name of the configuration
	Name string `json:"name" yaml:"name"`
	// TendermintRPC is the RPC to use when talking with tendermint
	TendermintRPC string `json:"tendermint_rpc" yaml:"tendermintRPC"`
	// APIServerAddr is the connection address towards the APIServer
	APIServerAddr    string `json:"api_server_addr" yaml:"apiServerAddr"`
	KeyringDirectory string `json:"keyring_directory" yaml:"keyringDirectory"`
	KeyringBackend   string `json:"keyring_backend" yaml:"keyringBackend"`
	DefaultAccount   string `json:"default_account" yaml:"defaultAccount"`
}

func (c *Config) Validate() error {
	if c.TendermintRPC == "" || c.APIServerAddr == "" {
		return fmt.Errorf("missing tendermint and apiserver addresses (one at least required)")
	}
	if c.KeyringDirectory == "" {
		klog.Warningf("missing keyring directory in config")
	}
	return nil
}

func GetConfig() (*Config, error) {
	// first check if env is set
	path, ok := os.LookupEnv(ConfigDirectoryEnv)
	if !ok {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		path = filepath.Join(homeDir, DefaultConfigDirectory)
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
