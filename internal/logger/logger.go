package logger

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"ostkost/go-ps-hw-fiber/config"
	"path/filepath"
	"strings"
)

type LoggerOutput interface {
	io.Writer
}

func NewLogger(config *config.LoggerConfig) *slog.Logger {
	slog.SetLogLoggerLevel(parseLogLevel(config.Level))
	var logger *slog.Logger
	output, err := NewOutput(config)
	if err != nil {
		fmt.Println("Error creating logger:", err)
	}
	if config.Format == "json" {
		logger = slog.New(slog.NewJSONHandler(output, nil))
	} else {
		logger = slog.New(slog.NewTextHandler(output, nil))
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

func NewOutput(config *config.LoggerConfig) (LoggerOutput, error) {
	switch config.Type {
	case "stdout":
		return os.Stdout, nil
	case "stderr":
		return os.Stderr, nil
	case "file":
		if config.FilePath == "" {
			return os.Stdout, errors.New("file path not specified")
		}
		// Проверяем, что директория существует или может быть создана
		dir := filepath.Dir(config.FilePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create log directory: %w", err)
		}
		// Проверяем, что путь доступен для записи
		if err := verifyWritable(dir); err != nil {
			return nil, fmt.Errorf("log directory is not writable: %w", err)
		}
		return os.OpenFile(config.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	default:
		return os.Stdout, errors.New(fmt.Sprintf("unknown output type: %s", config.Type))
	}
}

func verifyWritable(dir string) error {
	// Пытаемся создать тестовый файл
	testFile := filepath.Join(dir, ".testwrite")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		return err
	}
	// Удаляем тестовый файл
	_ = os.Remove(testFile)
	return nil
}
