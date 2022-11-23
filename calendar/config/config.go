package config

import (
	"context"
	"fmt"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"github.com/heetch/confita/backend/flags"
	"os"
)

type Config struct {
	ConfigFile string `config:"config"`
	HttpListen string `config:"http_listen" yaml:"http_listen"`
	LogFile    string `config:"log_file" yaml:"log_file"`
	LogLevel   string `config:"log_level" yaml:"log_level"`
}

func ParseConfig() Config {
	var configFile string = "config.yml"

	cfg := Config{
		ConfigFile: configFile,
		HttpListen: "127.0.0.1:8080",
		LogFile:    "logs.txt",
		LogLevel:   "info",
	}

	for i, v := range os.Args[1:] {
		if v == "--config" || v == "-config" {
			configFile = os.Args[i+2]
		}
	}

	fullLoader := confita.NewLoader(
		env.NewBackend(),
		file.NewBackend(configFile),
		flags.NewBackend(),
	)

	if err := fullLoader.Load(context.Background(), &cfg); err != nil {
		fmt.Errorf("fail to read config %e", err)
		panic(1)
	}

	return cfg
}
