package clients

import (
	"gateway/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	clients "gateway/internal/genproto/clients"
	contructor "gateway/internal/genproto/contractors"

	auth "gateway/internal/genproto/auth"
)

type Clients struct {
	Client     clients.ClientServiceClient
	Contractor contructor.BidServiceClient
	User       auth.AuthServiceClient
}

func NewClients(cfg *config.Config) (*Clients, error) {
	path := cfg.MEDALS_HOST + cfg.MEDALS_PORT

	conn, err := grpc.NewClient(path, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := clients.NewClientServiceClient(conn)
	contractor := contructor.NewBidServiceClient(conn)
	user := auth.NewAuthServiceClient(conn)

	return &Clients{
		Client:     client,
		Contractor: contractor,
		User:       user,

	}, nil
}
