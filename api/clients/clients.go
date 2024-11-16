package clients

import (
	"gateway/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	user "gateway/internal/genproto/user-service"
	f "gateway/internal/genproto/first-service"
	s "gateway/internal/genproto/second-service"
	t "gateway/internal/genproto/third-service"
	fr "gateway/internal/genproto/fourth-service"
)

type Clients struct {
	User user.UserServiceClient
	First  f.FirstServiceClient
	Second s.SecondServiceClient
	Third  t.ThirdServiceClient
	Fourth fr.FourthServiceClient
}

func NewClients(cfg *config.Config) (*Clients, error) {
	path := cfg.MEDALS_HOST + cfg.MEDALS_PORT

	conn, err := grpc.NewClient(path, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	User1 := user.NewUserServiceClient(conn)
	First1 := f.NewFirstServiceClient(conn)
	Second1 := s.NewSecondServiceClient(conn)
	Third1:= t.NewThirdServiceClient(conn)
	Fourth1 := fr.NewFourthServiceClient(conn)

	return &Clients{
		User: User1,
		First: First1,
		Second: Second1,
		Third: Third1,
		Fourth: Fourth1,
	}, nil
}
