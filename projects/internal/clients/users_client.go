package clients

import (
	"context"
	"fmt"

	"github.com/LesterFernandes/tasks/shared-protos/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UsersServiceClient struct {
	Client pb.UsersServiceClient
}

func InitUsersServiceClient(url string) (*UsersServiceClient, error) {
	gc, err := grpc.NewClient(
		url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Error().Err(err).Msg("error connecting users grpc client")
		return nil, err
	}

	c := &UsersServiceClient{
		Client: pb.NewUsersServiceClient(gc),
	}

	return c, nil
}

func (c *UsersServiceClient) GetUsers() (*pb.ListUsersResponse, error) {
	users, err := c.Client.ListUsers(context.Background(), &pb.ListUsersRequest{})
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}
	fmt.Println(users)
	return users, nil
}
