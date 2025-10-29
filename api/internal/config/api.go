package config

import (
	"log"
)

type ApiConfig struct {
	DragonflyDB        int
	DragonflyHost      string
	DragonflyPassword  string
	DragonflyPort      string
	DragonflyProtocol  int
	Environment        string
	Host               string
	LogLevel           string
	Port               string
	QueueConcurrency   int
	QueueName          string
	WebsocketQueueName string
	ServerQueueName    string
}

func LoadApi() *ApiConfig {
	cfg := &ApiConfig{
		DragonflyDB:        getEnvAsInt("DRAGONFLY_DB", 0),
		DragonflyHost:      getEnv("DRAGONFLY_HOST", "localhost"),
		DragonflyPassword:  getEnv("DRAGONFLY_PASSWORD", ""),
		DragonflyPort:      getEnv("DRAGONFLY_PORT", "6379"),
		DragonflyProtocol:  getEnvAsInt("DRAGONFLY_PROTOCOL", 2),
		Environment:        getEnv("ENVIRONMENT", "development"),
		Host:               getEnv("HOST", "localhost"),
		LogLevel:           getEnv("LOG_LEVEL", "info"),
		Port:               getEnv("PORT", "8080"),
		QueueConcurrency:   getEnvAsInt("QUEUE_CONCURRENCY", 10),
		QueueName:          getEnv("QUEUE_NAME", "websocket-tasks"),
		WebsocketQueueName: getEnv("WEBSOCKET_QUEUE_NAME", "websocket-tasks"),
		ServerQueueName:    getEnv("SERVER_QUEUE_NAME", "server-tasks"),
	}

	log.Printf("API configuration loaded: %+v", cfg)

	return cfg
}
