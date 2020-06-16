package library

import (
	"os"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func init() {
	env := os.Getenv("ENV_NAME")
	if env == "prod" {
		config := zap.NewProductionConfig()
		config.Encoding = "console"
		logger, _ = config.Build()
	} else {
		logger, _ = zap.NewDevelopment()
	}
}

func Logger() *zap.Logger {
	return logger
}

func SugaredLogger() *zap.SugaredLogger {
	return logger.Sugar()
}
