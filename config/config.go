package config

import (
	"time"

	"github.com/spf13/viper"
)

var configs *Config
var GlobalConfig *Config

func LoadConfig(config *string) (*Config, error) {
	if configs != nil {
		return configs, nil
	}

	//Path to config file
	viper.AddConfigPath(".")
	viper.SetConfigName(*config)
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	//Default values
	viper.SetDefault("database.timeout", 10*time.Second)
	viper.SetDefault("database.dbname", "testing")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	configs := &Config{}
	if err := viper.Unmarshal(configs); err != nil {
		configs = nil
		return configs, err
	}

	GlobalConfig = configs
	return configs, nil
}
