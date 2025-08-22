package services

import (
	db "github.com/LesterFernandes/tasks/users/db/gen"
	"github.com/LesterFernandes/tasks/shared-protos/pb"
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
