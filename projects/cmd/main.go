package main

import (
	"github.com/LesterFernandes/tasks/projects/internal/clients"
	"github.com/rs/zerolog/log"
)

func main() {
	c, err := clients.InitUsersServiceClient("localhost:9001")
	if err != nil {
		log.Error().Err(err)
		return
	}
	c.GetUsers()
}
