package config

import "github.com/spf13/viper"

var CONF = &BaseConfig{}

type BaseConfig struct {
	ListenAddr string
	Debug      bool
}

// InitConfiguration ...
func InitConfiguration(configName string, configPaths []string, config interface{}) error {
	vp := viper.New()
	vp.SetConfigName(configName)
	vp.AutomaticEnv()
	for _, configPath := range configPaths {
		vp.AddConfigPath(configPath)
	}

	if err := vp.ReadInConfig(); err != nil {
		return err
	}

	err := vp.Unmarshal(config)
	if err != nil {
		return err
	}

	return nil
}
