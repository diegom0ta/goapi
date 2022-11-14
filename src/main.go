package main

import (
	"log"
	"net/http"
)

const (
	port       = ":8090"
	reposRoute = "/api/repos"
	userRoute  = "/api/user"
	usersRoute = "/api/users"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(reposRoute, handlers.ReposHandler)
	mux.HandleFunc(userRoute, handlers.UserHandler)
	mux.HandleFunc(usersRoute, handlers.UsersHandler)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
