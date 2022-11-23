package logger

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
)

type Logger struct {
	zap.SugaredLogger
}

type Config struct {
	LogLevel string
	LogFile  string
}

func InitLogger(cfg Config) *zap.SugaredLogger {
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
	zapLogger := zap.Must(zcfg.Build())
	return zapLogger.Sugar()
}
