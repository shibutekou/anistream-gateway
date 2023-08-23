package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Env                string `yaml:"env" env-default:"local"`
	ProxyAddr          string `yaml:"proxy_addr" env-default:":50050"`
	ContentServiceAddr string `yaml:"content_service_addr" env-required:"true"`
}

func Load() *Config {
	configPath := os.Getenv("GATEWAY_SERVICE_CONFIG_PATH")
	if configPath == "" {
		log.Fatal("GATEWAY_SERVICE_CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
