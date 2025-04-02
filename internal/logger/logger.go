package logger

import (
	"log/slog"
	"os"
	"ostkost/go-ps-hw-fiber/config"
	"strings"
)

func NewLogger(config *config.LoggerConfig) *slog.Logger {
	slog.SetLogLoggerLevel(parseLogLevel(config.Level))
	var logger *slog.Logger
	if config.Format == "json" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	return logger
}

func parseLogLevel(levelStr string) slog.Level {
	switch strings.ToUpper(levelStr) {
	case "DEBUG", "7":
		return slog.LevelDebug
	case "INFO", "6", "5":
		return slog.LevelInfo
	case "WARN", "WARNING", "4":
		return slog.LevelWarn
	case "ERROR", "3", "2", "1", "0":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
