package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type DbConfig struct {
	ConnectionUri string `mapstructure:"connection_uri"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type AuthConfig struct {
	Ttl    time.Duration `mapstructure:"ttl"`
	Secret string        `mapstructure:"secret"`
}

type OAuthConfig struct {
	GoogleApiUrl string `mapstructure:"google_api_url"`
}

type AppConfig struct {
	DbConfig     DbConfig     `mapstructure:"db"`
	ServerConfig ServerConfig `mapstructure:"server"`
	AuthConfig   AuthConfig   `mapstructure:"auth"`
	OAuthConfig  OAuthConfig  `mapstructure:"oauth"`
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
