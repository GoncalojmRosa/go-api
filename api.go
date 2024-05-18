package main

import (
	"fmt"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

type User struct {
	ID string
	Username string
	Email string
}

// constructor
func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr}
}


func handleGetUser(w http.ResponseWriter, r *http.Request){
	userID := r.PathValue("userID")

	// simulate fetching user from database
	user := &User{
		ID: userID,
		Username: "example",
		Email: "example@gmail.com",
	}

	w.Write([]byte(fmt.Sprintf("User: %v\n", user)))
}

func (s *APIServer) Start() error {
	router := http.NewServeMux()

	router.HandleFunc("GET /users/{userID}", handleGetUser)

	server := http.Server{
		Addr: s.addr,
		Handler: router,
	}

	log.Printf("API server listening on %s\n", s.addr)

	return server.ListenAndServe()
}