package handlers

import (
	"log"
	"net/http"

	apiParser "github.com/diegom0ta/goapi/src/parser"
)

const (
	repos = "https://api.github.com/users/diegom0ta/repos"
	user  = "https://api.github.com/users/diegom0ta"
	users = "https://api.github.com/users"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Method: GET - Endpoints: /api/repos, /api/user, /api/users\n"))
}

func ReposHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(repos)
	if err != nil {
		log.Fatal(err)
	}
	wr := apiParser.WriteResponse(resp.Body)
	w.Write([]byte(wr))
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(user)
	if err != nil {
		log.Fatal(err)
	}
	wr := apiParser.WriteResponse(resp.Body)
	w.Write([]byte(wr))
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(users)
	if err != nil {
		log.Fatal(err)
	}
	wr := apiParser.WriteResponse(resp.Body)
	w.Write([]byte(wr))
}
