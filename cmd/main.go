package main

import (
	"gateway/api"
	"gateway/api/handler"
	"gateway/internal/config"
	"gateway/internal/logs"

	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()

	log, err := logs.NewLogger()
	if err != nil {
		return
	}

	handler, err := handler.NewHandler(&cfg)
	if err != nil {
		log.Error("Error in main")
	}

	r := api.InitRouter(handler)

	log.Info("Starting server on port :%s", zap.Any("", cfg.HTTP_PORT))
	if err = r.Run(cfg.HTTP_PORT); err != nil {
		log.Fatal("error starting server ... ")
	}

}