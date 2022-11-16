package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"github.com/heetch/confita/backend/flags"
	"go.uber.org/zap"
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

func InitLogger(cfg Config) *zap.Logger {
	rawJSON := []byte(fmt.Sprintf(`{
		"level": "%s",
		"encoding": "json",
		"outputPaths": ["stdout", "%s"],
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase",
			"timeKey": "timestamp",
			"timeEncoder": "iso8601"
		}
	  }`, cfg.LogLevel, cfg.LogFile))

	var zcfg zap.Config
	if err := json.Unmarshal(rawJSON, &zcfg); err != nil {
		panic(err)
	}
	return zap.Must(zcfg.Build())
}

func main() {
	cfg := ParseConfig()

	logger := InitLogger(cfg)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		logger.Debug(fmt.Sprintf("GET /hello %s", r.RemoteAddr))
		fmt.Fprintln(w, "hello")
	})

	logger.Info(fmt.Sprintf("Server starting at %s", cfg.HttpListen))
	http.ListenAndServe(cfg.HttpListen, nil)
}
