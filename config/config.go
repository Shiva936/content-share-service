package config

import (
	"os"
)

type Config struct {
	AppName          string `json:"app_name"`
	URL              string `json:"url"`
	Port             string `json:"port"`
	LogLevel         string `json:"log_level"`
	DatabaseURL      string `json:"database_url"`
	DatabaseInstance string `json:"database_instance"`
}

var conf *Config

func Get() *Config {
	return conf
}

func Set(c *Config) {
	conf = c
	conf.DatabaseURL = os.Getenv("DATABASE_URL")
	conf.DatabaseInstance = os.Getenv("DATABASE_INSTANCE")
}
