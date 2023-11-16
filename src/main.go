package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	h "github.com/diegom0ta/goapi/src/handlers"
)

var port = os.Getenv("PORT")

const (
	homeRoute  = "/"
	reposRoute = "/api/repos"
	userRoute  = "/api/user"
	usersRoute = "/api/users"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(homeRoute, h.HomeHandler)
	mux.HandleFunc(reposRoute, h.ReposHandler)
	mux.HandleFunc(userRoute, h.UserHandler)
	mux.HandleFunc(usersRoute, h.UsersHandler)

	if port == "" {
		port = "9000"
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Server listening on port: %s\n", port)

	go func() {
		if err := srv.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Listen and serve error: %v\n", err)
		}
	}()

	<-ctx.Done()

	log.Println("Got interruption signal")

	if err := srv.Shutdown(context.TODO()); err != nil {
		log.Printf("Server shutdown returned an error: %v\n", err)
	}
}
