package clients

import (
	"github.com/LesterFernandes/tasks/users/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UsersServiceClient struct {
	Client pb.UsersServiceClient
}

func InitUsersServiceClient(url string) (*UsersServiceClient, error) {
	cc, err := grpc.NewClient(
		url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Error().Err(err).Msg("error connecting users grpc client")
		return nil, err
	}

	c := &UsersServiceClient{
		Client: pb.NewUsersServiceClient(cc),
	}

	return c, nil
}

// func (c *UsersServiceClient) GetUsers() (*pb.ListUsersResponse, error) {
// 	c.Client.CreateUser()
// }
