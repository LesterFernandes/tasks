package clients

import (
	"net/http"

	pb "github.com/LesterFernandes/tasks/shared-protos/pb"
	"github.com/LesterFernandes/tasks/utils"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UsersServiceClient struct {
	Client pb.UsersServiceClient
}

func InitUsersServiceClient(url string) (*UsersServiceClient, error) {
	gc, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	return &UsersServiceClient{
		Client: pb.NewUsersServiceClient(gc),
	}, nil
}

func (c *UsersServiceClient) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := utils.ReadBody[pb.CreateUserRequest](r)

	if err != nil {
		utils.Resp(w, http.StatusInternalServerError, "")
		return
	}

	resp, err := c.Client.CreateUser(r.Context(), body)
	if err != nil {
		utils.Resp(w, http.StatusInternalServerError, "")
		return
	}

	utils.Resp(w, http.StatusOK, resp)
}

func (c *UsersServiceClient) ListUsers(w http.ResponseWriter, r *http.Request) {
	list, err := c.Client.ListUsers(r.Context(), &pb.ListUsersRequest{})
	if err != nil {
		utils.Resp(w, http.StatusInternalServerError, "")
		return
	}
	utils.Resp(w, http.StatusOK, list)
}
