package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		URL string `mapstructure:"url"`
	} `mapstructure:"database"`
	Frontend struct {
		Dev  string
		Dev2 string
		Prod string
	}
	AppEnv string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AutomaticEnv()

	var cfg Config

	if err := viper.ReadInConfig(); err == nil {
		if err := viper.Unmarshal(&cfg); err != nil {
			return nil, err
		}
	} else {
		cfg.Server.Port = getenv("APP_PORT", "3000")
		cfg.Database.URL = os.Getenv("DATABASE_URL")
		cfg.AppEnv = getenv("APP_ENV", "development")
		cfg.Frontend.Dev = os.Getenv("FRONTEND_URL_DEV")
		cfg.Frontend.Dev2 = os.Getenv("FRONTEND_URL_DEV2")
		cfg.Frontend.Prod = os.Getenv("FRONTEND_URL_PROD")
	}

	if cfg.Database.URL == "" {
		return nil, errors.New("DATABASE_URL is required")
	}

	return &cfg, nil
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
