package logger

import (
	"log/slog"
	"testing"
)

func TestParseLogLevel(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected slog.Level
	}{
		{"Debug level - DEBUG", "DEBUG", slog.LevelDebug},
		{"Debug level - 7", "7", slog.LevelDebug},
		{"Info level - INFO", "INFO", slog.LevelInfo},
		{"Info level - 6", "6", slog.LevelInfo},
		{"Info level - 5", "5", slog.LevelInfo},
		{"Warn level - WARN", "WARN", slog.LevelWarn},
		{"Warn level - WARNING", "WARNING", slog.LevelWarn},
		{"Warn level - 4", "4", slog.LevelWarn},
		{"Error level - ERROR", "ERROR", slog.LevelError},
		{"Error level - 3", "3", slog.LevelError},
		{"Error level - 2", "2", slog.LevelError},
		{"Error level - 1", "1", slog.LevelError},
		{"Error level - 0", "0", slog.LevelError},
		{"Default case - empty string", "", slog.LevelInfo},
		{"Default case - unknown level", "CRITICAL", slog.LevelInfo},
		{"Case insensitive - debug", "debug", slog.LevelDebug},
		{"Case insensitive - warn", "warn", slog.LevelWarn},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseLogLevel(tt.input)
			if got != tt.expected {
				t.Errorf("parseLogLevel(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
