package main

import (
	"net/http"

	"github.com/LesterFernandes/tasks/api-gateway/internal/clients"
	"github.com/LesterFernandes/tasks/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	usersClient, err := clients.InitUsersServiceClient("localhost:9001")
	if err != nil {
		return
	}

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		utils.Resp(w, http.StatusOK, "PONG")
	})

	r.Route("/users", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Get("/", usersClient.ListUsers)
		r.Post("/create", usersClient.CreateUser)
	})

	http.ListenAndServe(":3000", r)
}
