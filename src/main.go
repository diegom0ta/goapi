package main

import (
	"log"
	"net/http"
	"os"

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
		port = ":5000"
	}
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
