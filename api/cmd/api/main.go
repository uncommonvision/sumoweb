package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"sumoweb/server/internal/config"
	"sumoweb/server/internal/queue"
	"sumoweb/server/internal/routes"
	websocketWorker "sumoweb/server/internal/workers/websocket"
)

func main() {
	cfg := config.LoadApi()

	_, err := queue.InitClient(cfg)
	if err != nil {
		log.Printf("Warning: Failed to connect to Dragonfly: %v", err)
		log.Println("Server will continue without queue support")
	}
	defer queue.Close()

	_, err = queue.InitProducer(cfg)
	if err != nil {
		log.Printf("Warning: Failed to initialize queue producer: %v", err)
	}
	defer func() {
		if p := queue.GetProducer(); p != nil {
			p.Close()
		}
	}()

	workerCfg := &config.WorkerConfig{
		DragonflyDB:        cfg.DragonflyDB,
		DragonflyHost:      cfg.DragonflyHost,
		DragonflyPassword:  cfg.DragonflyPassword,
		DragonflyPort:      cfg.DragonflyPort,
		DragonflyProtocol:  cfg.DragonflyProtocol,
		QueueConcurrency:   1,
		QueueName:          cfg.WebsocketQueueName,
		WebsocketQueueName: cfg.WebsocketQueueName,
		ServerQueueName:    cfg.ServerQueueName,
	}

	broadcastWorker, err := queue.NewWorker(workerCfg)
	if err != nil {
		log.Printf("Warning: Failed to create broadcast worker: %v", err)
	} else {
		broadcastWorker.RegisterHandler(queue.TypeBroadcastMatches, websocketWorker.HandleBroadcastMatches)

		go func() {
			if err := broadcastWorker.Start(); err != nil {
				log.Printf("Broadcast worker error: %v", err)
			}
		}()
		defer broadcastWorker.Shutdown()

		log.Println("Broadcast worker started in API process")
	}

	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	routes.Setup(router)

	srv := &http.Server{
		Addr:         cfg.Host + ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	go func() {
		log.Printf("Server starting on %s:%s", cfg.Host, cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
