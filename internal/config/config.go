package config

import (
	"log"

	"github.com/spf13/viper"
)

type DbConfig struct {
	ConnectionUri string `mapstructure:"connection_uri"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type AppConfig struct {
	DbConfig     DbConfig     `mapstructure:"db"`
	ServerConfig ServerConfig `mapstructure:"server"`
}

var Config AppConfig

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Error unmarshaling envs: %s", err)
	}
}
