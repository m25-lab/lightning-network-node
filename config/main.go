package config

import (
	"time"

	"github.com/spf13/viper"
)

func LoadConfig() (Config, error) {
	//Path to config file
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	//Default values
	viper.SetDefault("database.timeout", 10*time.Second)

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var configs Config
	err := viper.Unmarshal(&configs)
	if err != nil {
		return Config{}, err
	}

	return configs, nil
}
