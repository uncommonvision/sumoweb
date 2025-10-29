package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"sumoweb/server/internal/config"
	"sumoweb/server/internal/queue"
	matchWorker "sumoweb/server/internal/workers/match"
	rikishiWorker "sumoweb/server/internal/workers/rikishi"
)

func main() {
	cfg := config.LoadWorker()

	log.Printf("Worker starting: queue=%s concurrency=%d", cfg.QueueName, cfg.QueueConcurrency)

	apiCfg := &config.ApiConfig{
		DragonflyDB:        cfg.DragonflyDB,
		DragonflyHost:      cfg.DragonflyHost,
		DragonflyPassword:  cfg.DragonflyPassword,
		DragonflyPort:      cfg.DragonflyPort,
		DragonflyProtocol:  cfg.DragonflyProtocol,
		QueueName:          cfg.ServerQueueName,
		WebsocketQueueName: cfg.WebsocketQueueName,
		ServerQueueName:    cfg.ServerQueueName,
	}

	_, err := queue.InitClient(apiCfg)
	if err != nil {
		log.Fatalf("Failed to connect to Dragonfly: %v", err)
	}
	defer queue.Close()

	_, err = queue.InitProducer(apiCfg)
	if err != nil {
		log.Printf("Warning: Failed to initialize queue producer: %v", err)
	}
	defer func() {
		if p := queue.GetProducer(); p != nil {
			p.Close()
		}
	}()

	worker, err := queue.NewWorker(cfg)
	if err != nil {
		log.Fatalf("Failed to create worker: %v", err)
	}
	defer worker.Shutdown()

	worker.RegisterHandler(queue.TypeFetchMatches, matchWorker.HandleFetchMatches)
	worker.RegisterHandler(queue.TypeFetchRikishiProfile, rikishiWorker.HandleFetchRikishiProfile)

	go func() {
		if err := worker.Start(); err != nil {
			log.Fatalf("Worker failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Worker exiting gracefully")
}
