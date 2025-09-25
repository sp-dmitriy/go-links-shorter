package main

import (
	"fmt"
	"go-links-shorter/configs"
	"go-links-shorter/internal/auth"

	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
