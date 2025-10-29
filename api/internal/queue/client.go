package queue

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"sumoweb/server/internal/config"
)

var redisClient *redis.Client

func InitClient(cfg *config.ApiConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.DragonflyHost, cfg.DragonflyPort),
		DB:       cfg.DragonflyDB,
		Password: cfg.DragonflyPassword,
		Protocol: cfg.DragonflyProtocol,
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Dragonfly: %w", err)
	}

	redisClient = client
	log.Printf("Connected to Dragonfly at %s:%s (protocol: RESP%d)", cfg.DragonflyHost, cfg.DragonflyPort, cfg.DragonflyProtocol)
	return client, nil
}

func GetClient() *redis.Client {
	return redisClient
}

func Close() error {
	if redisClient != nil {
		return redisClient.Close()
	}
	return nil
}
