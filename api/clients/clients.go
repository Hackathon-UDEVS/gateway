package clients

import (
	"fmt"
	"github.com/Hackaton-UDEVS/gateway/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	clients "github.com/Hackaton-UDEVS/gateway/internal/genproto/auth"
	contructor "github.com/Hackaton-UDEVS/gateway/internal/genproto/tender-service"

	auth "github.com/Hackaton-UDEVS/gateway/internal/genproto/auth"
)

type Clients struct {
	Client     contructor.ClientServiceClient
	Contractor contructor.ContractorServiceClient
	User       clients.AuthServiceClient
}

func NewClients(cfg *config.Config) (*Clients, error) {
	authClient := fmt.Sprintf("%s:%d", cfg.AUTHHOST, cfg.AUTHPORT)
	fmt.Println(authClient)
	conn, err := grpc.NewClient(authClient, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	conn1, err := grpc.NewClient(authClient, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := contructor.NewClientServiceClient(conn)
	contractor := contructor.NewContractorServiceClient(conn)
	user := auth.NewAuthServiceClient(conn1)

	return &Clients{
		Client:     client,
		Contractor: contractor,
		User:       user,
	}, nil
}
