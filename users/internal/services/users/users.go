package services

import (
	"context"

	"github.com/LesterFernandes/tasks/shared-protos/pb"
	db "github.com/LesterFernandes/tasks/users/db/gen"
	"github.com/rs/zerolog/log"
)

type service struct {
	db *db.Queries
	pb.UnimplementedUsersServiceServer
}

func NewService(dbConn *db.Queries) *service {
	return &service{
		db: dbConn,
	}
}

func (s *service) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		log.Error().Err(err).Msg("DB error retrieving users")
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

func (s *service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	resp, err := s.db.CreateUser(ctx, db.CreateUserParams{
		PasswordHash: "_pw",
		Email:        req.Email,
		UserName:     req.UserName,
		Name:         req.Name,
	})

	if err != nil {
		log.Error().Err(err).Msg("Error creating user")
		return nil, err
	}

	return &pb.CreateUserResponse{
		Status: 200,
		UserId: resp.UserID.String(),
	}, nil
}

// func (s *service) CreateTeam(ctx context.Context, req *pb.CreateTeamRequest) (*pb.CreateTeamResponse, error) {
// }
