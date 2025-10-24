package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DragonflyDB       int
	DragonflyHost     string
	DragonflyPassword string
	DragonflyPort     string
	DragonflyProtocol int
	Environment       string
	Host              string
	LogLevel          string
	Port              string
	QueueConcurrency  int
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		DragonflyDB:       getEnvAsInt("DRAGONFLY_DB", 0),
		DragonflyHost:     getEnv("DRAGONFLY_HOST", "localhost"),
		DragonflyPassword: getEnv("DRAGONFLY_PASSWORD", ""),
		DragonflyPort:     getEnv("DRAGONFLY_PORT", "6379"),
		DragonflyProtocol: getEnvAsInt("DRAGONFLY_PROTOCOL", 2),
		Environment:       getEnv("ENVIRONMENT", "development"),
		Host:              getEnv("HOST", "localhost"),
		LogLevel:          getEnv("LOG_LEVEL", "info"),
		Port:              getEnv("PORT", "8080"),
		QueueConcurrency:  getEnvAsInt("QUEUE_CONCURRENCY", 10),
	}

	log.Printf("Configuration loaded: environment=%s host=%s port=%s", cfg.Environment, cfg.Host, cfg.Port)
	return cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
