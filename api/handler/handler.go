package handler

import (
	"gateway/api/clients"
	"gateway/api/kafka"
	"gateway/internal/config"
)

type Handler struct {
	Clients  *clients.Clients
	Producer kafka.KafkaProducer
}

func NewHandler(cfg *config.Config) (*Handler, error) {
	clients, err := clients.NewClients(cfg)
	if err != nil {
		return nil, err
	}

	kafkaProd, err := kafka.NewKafkaProducer([]string{cfg.KAFKA_HOST + cfg.KAFKA_PORT})
	if err != nil {
		return nil, err
	}

	return &Handler{
		Clients:  clients,
		Producer: kafkaProd,
	}, nil

}
