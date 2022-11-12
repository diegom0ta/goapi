package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

var srv = new(http.Server)

const (
	repos = "https://api.github.com/users/diegom0ta/repos"
	user  = "https://api.github.com/users/diegom0ta"
	users = "https://api.github.com/users"
)

func getRepos() io.ReadCloser {
	resp, err := http.Get(repos)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Body
}

func getUser() io.ReadCloser {
	resp, err := http.Get(user)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Body
}

func getUsers() io.ReadCloser {
	resp, err := http.Get(users)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Body
}

func writeResponse(rc io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(rc)
	respBytes := buf.String()

	return string(respBytes)
}

func main() {
	listRepos := getRepos()
	fmt.Print(writeResponse(listRepos))

	usr := getUser()
	fmt.Print(writeResponse(usr))

	listUsers := getUsers()
	fmt.Print(writeResponse(listUsers))

}
