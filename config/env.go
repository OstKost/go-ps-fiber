package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type DatabaseConfig struct {
	dbUrl string
}

type LoggerConfig struct {
	Format   string
	Level    string
	Type     string
	FilePath string
}

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not loaded")
		return
	}
	log.Println(".env loaded successfully")
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		dbUrl: getString("DATABASE_URL", ""),
	}
}

func NewLoggerConfig() *LoggerConfig {
	return &LoggerConfig{
		Format:   getString("LOG_FORMAT", "text"),
		Level:    getString("LOG_LEVEL", "info"),
		Type:     getString("LOG_TYPE", "stdout"),
		FilePath: getString("LOG_FILE_PATH", ""),
	}
}

func getString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	i, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return i
}

func getInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return i
}
