package queue

import (
	"context"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
	"sumoweb/server/internal/config"
)

type Worker struct {
	mux    *asynq.ServeMux
	server *asynq.Server
}

func NewWorker(cfg *config.WorkerConfig) (*Worker, error) {
	redisOpt := asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%s", cfg.DragonflyHost, cfg.DragonflyPort),
		DB:       cfg.DragonflyDB,
		Password: cfg.DragonflyPassword,
	}

	queues := map[string]int{
		"default": 1,
	}

	if cfg.QueueName != "" && cfg.QueueName != "default" {
		queues[cfg.QueueName] = 10
	}

	server := asynq.NewServer(redisOpt, asynq.Config{
		Concurrency: cfg.QueueConcurrency,
		Queues:      queues,
	})

	mux := asynq.NewServeMux()

	log.Printf("Worker configured: %+v", cfg)

	return &Worker{
		mux:    mux,
		server: server,
	}, nil
}

func (w *Worker) RegisterHandler(pattern string, handler func(context.Context, *asynq.Task) error) {
	w.mux.HandleFunc(pattern, handler)
}

func (w *Worker) Shutdown() {
	log.Println("Shutting down worker...")
	if w.server != nil {
		w.server.Shutdown()
	}
}

func (w *Worker) Start() error {
	log.Println("Starting worker...")
	return w.server.Run(w.mux)
}
