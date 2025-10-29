package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
	"sumoweb/server/internal/config"
	"sumoweb/server/internal/models"
)

var producer *Producer

type Producer struct {
	client             *asynq.Client
	websocketQueueName string
	serverQueueName    string
}

func InitProducer(cfg *config.ApiConfig) (*Producer, error) {
	redisOpt := asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%s", cfg.DragonflyHost, cfg.DragonflyPort),
		DB:       cfg.DragonflyDB,
		Password: cfg.DragonflyPassword,
	}

	client := asynq.NewClient(redisOpt)

	producer = &Producer{
		client:             client,
		websocketQueueName: cfg.WebsocketQueueName,
		serverQueueName:    cfg.ServerQueueName,
	}

	log.Printf("Queue producer initialized: %+v", cfg)
	return producer, nil
}

func GetProducer() *Producer {
	return producer
}

func (p *Producer) Close() error {
	if p.client != nil {
		return p.client.Close()
	}
	return nil
}

func (p *Producer) EnqueueFetchMatches(ctx context.Context, division, day int) (string, error) {
	payload, err := json.Marshal(FetchMatchesPayload{
		Day:      day,
		Division: division,
	})
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	task := asynq.NewTask(TypeFetchMatches, payload)

	info, err := p.client.Enqueue(task, asynq.Queue(p.serverQueueName))
	if err != nil {
		return "", fmt.Errorf("failed to enqueue task: %w", err)
	}

	log.Printf("Enqueued task: id=%s type=%s queue=%s", info.ID, info.Type, info.Queue)
	return info.ID, nil
}

func (p *Producer) EnqueueFetchRikishiProfile(ctx context.Context, rikishiID string) (string, error) {
	payload, err := json.Marshal(FetchRikishiProfilePayload{
		RikishiID: rikishiID,
	})
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	task := asynq.NewTask(TypeFetchRikishiProfile, payload)

	info, err := p.client.Enqueue(task, asynq.Queue(p.serverQueueName))
	if err != nil {
		return "", fmt.Errorf("failed to enqueue task: %w", err)
	}

	log.Printf("Enqueued task: id=%s type=%s queue=%s", info.ID, info.Type, info.Queue)
	return info.ID, nil
}

func (p *Producer) EnqueueBroadcastMatches(ctx context.Context, division, day int, matches []models.Match) (string, error) {
	payload, err := json.Marshal(BroadcastMatchesPayload{
		Division: division,
		Day:      day,
		Matches:  matches,
	})
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	task := asynq.NewTask(TypeBroadcastMatches, payload)

	info, err := p.client.Enqueue(task, asynq.Queue(p.websocketQueueName))
	if err != nil {
		return "", fmt.Errorf("failed to enqueue task: %w", err)
	}

	log.Printf("Enqueued task: id=%s type=%s queue=%s", info.ID, info.Type, info.Queue)
	return info.ID, nil
}
