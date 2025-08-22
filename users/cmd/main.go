package main

import (
	"context"
	"net"
	"os"

	db "github.com/LesterFernandes/tasks/users/db/gen"
	"github.com/LesterFernandes/tasks/users/internal/services"
	"github.com/LesterFernandes/tasks/users/pb"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"

	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()
	// Setup database connection pool
	pool, err := pgxpool.New(ctx, os.Getenv("postgres://postgres:admin@localhost:5432/tasks"))
	if err != nil {
		log.Error().Err(err).Msg("Error db init")
	}
	defer pool.Close()

	dbConn := db.New(pool)
	service := services.NewService(dbConn)

	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Error().Err(err).Msg("Tcp conn error")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUsersServiceServer(grpcServer, service)

	if err := grpcServer.Serve(lis); err != nil {
		log.Error().Err(err).Msg("grpc server conn error")
	}
}
