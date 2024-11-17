package main

import (
	"fmt"
	"github.com/Hackaton-UDEVS/gateway/api"
	"github.com/Hackaton-UDEVS/gateway/api/handler"
	"github.com/Hackaton-UDEVS/gateway/internal/config"
	"github.com/Hackaton-UDEVS/gateway/internal/logs"
	"go.uber.org/zap"
	"os"

	_ "github.com/Hackaton-UDEVS/gateway/docs"
)

func main() {
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

	// Log the DB password for debugging (NOT RECOMMENDED IN PRODUCTION)
	fmt.Println("DB Password:", cfg.DBPASSWORD)

	// Start the server with proper error handling
	serverAddress := fmt.Sprintf("%s:%d", cfg.GATEWAYHOST, cfg.GATEWAYPORT)
	log.Info("Starting server", zap.String("address", serverAddress))

	// Gracefully start the server
	if err := r.Run(serverAddress); err != nil {
		log.Fatal("Failed to start server", zap.Error(err))
		os.Exit(1) // Ensure process exits if server fails to start
	}
}
