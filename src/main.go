package main

import (
	"log"
	"net/http"

	h "github.com/diegom0ta/goapi/src/handlers"
)

const (
	port       = ":8090"
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

	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
