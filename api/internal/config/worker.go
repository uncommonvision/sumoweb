package config

import (
	"log"
)

type WorkerConfig struct {
	DragonflyDB        int
	DragonflyHost      string
	DragonflyPassword  string
	DragonflyPort      string
	DragonflyProtocol  int
	Environment        string
	LogLevel           string
	QueueConcurrency   int
	QueueName          string
	WebsocketQueueName string
	ServerQueueName    string
}

func LoadWorker() *WorkerConfig {
	cfg := &WorkerConfig{
		DragonflyDB:        getEnvAsInt("DRAGONFLY_DB", 0),
		DragonflyHost:      getEnv("DRAGONFLY_HOST", "localhost"),
		DragonflyPassword:  getEnv("DRAGONFLY_PASSWORD", ""),
		DragonflyPort:      getEnv("DRAGONFLY_PORT", "6379"),
		DragonflyProtocol:  getEnvAsInt("DRAGONFLY_PROTOCOL", 2),
		Environment:        getEnv("ENVIRONMENT", "development"),
		LogLevel:           getEnv("LOG_LEVEL", "info"),
		QueueConcurrency:   getEnvAsInt("QUEUE_CONCURRENCY", 10),
		QueueName:          getEnv("QUEUE_NAME", "server-tasks"),
		WebsocketQueueName: getEnv("WEBSOCKET_QUEUE_NAME", "websocket-tasks"),
		ServerQueueName:    getEnv("SERVER_QUEUE_NAME", "server-tasks"),
	}

	log.Printf("Worker configuration loaded: %+v", cfg)

	return cfg
}
