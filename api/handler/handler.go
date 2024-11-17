package handler

import (
	"github.com/Hackaton-UDEVS/gateway/api/clients"
	"github.com/Hackaton-UDEVS/gateway/internal/config"
)

type Handler struct {
	Clients *clients.Clients
}

func NewHandler(cfg *config.Config) (*Handler, error) {

	client, err := clients.NewClients(cfg)
	if err != nil {
		return nil, err
	}

	return &Handler{
		Clients: client,
	}, nil
}
