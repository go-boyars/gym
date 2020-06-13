package config

import (
	"context"
	"log"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
)

type Config struct {
	Addr   string `config:"addr"`
	DBconn string `config:"dbconn"`
}

func GetConfig(configPath string) *Config {
	if configPath == "" {
		log.Fatal("no config file")
	}

	cfg := &Config{}

	loader := confita.NewLoader(
		file.NewBackend(configPath),
	)

	err := loader.Load(context.Background(), cfg)
	if err != nil {
		log.Fatal("cannot read config. ", err)
	}

	return cfg
}
