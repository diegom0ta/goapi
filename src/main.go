package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

const (
	repos      = "https://api.github.com/users/diegom0ta/repos"
	user       = "https://api.github.com/users/diegom0ta"
	users      = "https://api.github.com/users"
	port       = ":8090"
	reposRoute = "/api/repos"
	userRoute  = "/api/user"
	usersRoute = "/api/users"
)

func writeResponse(rc io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(rc)
	respBytes := buf.String()

	return string(respBytes)
}

func reposHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(repos)
	if err != nil {
		log.Fatal(err)
	}
	wr := writeResponse(resp.Body)
	log.Print(w.Write([]byte(wr)))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(user)
	if err != nil {
		log.Fatal(err)
	}
	wr := writeResponse(resp.Body)
	log.Print(w.Write([]byte(wr)))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(users)
	if err != nil {
		log.Fatal(err)
	}
	wr := writeResponse(resp.Body)
	log.Print(w.Write([]byte(wr)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(reposRoute, reposHandler)
	mux.HandleFunc(userRoute, userHandler)
	mux.HandleFunc(usersRoute, usersHandler)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
