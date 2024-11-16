package main

import (
	"fmt"
	"gateway/api"
	"gateway/api/handler"
	"gateway/internal/config"
	"gateway/internal/logs"
	"go.uber.org/zap"

	_ "gateway/docs"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize logger
	log, err := logs.NewLogger()
	if err != nil {
		// Handle logger initialization failure
		zap.S().Fatal("Failed to initialize logger:", err)
		return
	}

	// Initialize handlers
	handlers, err := handler.NewHandler(&cfg)
	if err != nil {
		log.Error("Error initializing handlers", zap.Error(err))
		return
	}

	// Initialize router
	r := api.InitRouter(handlers)
	fmt.Println(cfg.HTTP_PORT)
	// Start the server
	serverAddress := fmt.Sprintf(":%d", cfg.HTTP_PORT)
	log.Info("Starting server", zap.String("address", serverAddress))
	if err = r.Run(serverAddress); err != nil {
		log.Fatal("Failed to start server", zap.Error(err))
	}
}
