package services

import (
	"context"

	"github.com/LesterFernandes/tasks/users/pb"
	"github.com/rs/zerolog/log"
)

func (s *service) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {

	users, err := s.db.GetUsers(ctx)
	if err != nil {
		log.Error().Err(err).Msg("db error retrieving users")
		return nil, err
	}

	var usersResponse []*pb.UserResponse
	for _, user := range users {
		usersResponse = append(usersResponse, &pb.UserResponse{
			Email:    user.Email,
			UserName: user.UserName,
			Name:     user.Name,
			UserId:   user.UserID.String(),
		})
	}

	resp := &pb.ListUsersResponse{
		Status: 200,
		Users:  usersResponse,
	}

	return resp, nil
}
