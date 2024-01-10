package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if writeErr := WriteJSON(w, http.StatusBadRequest, err.Error()); writeErr != nil {
				fmt.Println("Error writing JSON:", writeErr)
			}
		}
	}
}

type APIServer struct {
	listenAddr string
}

func NewApiServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/endpoint", makeHTTPHandleFunc(s.handleEndpoint))

	log.Println("Json api server running on port: ", s.listenAddr)

	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func (s *APIServer) handleEndpoint(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		return s.handleGetEndpoint(w, r)
	}
	if r.Method == "POST" {
		return s.handlePostEndpoint(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetEndpoint(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handlePostEndpoint(w http.ResponseWriter, r *http.Request) error {
	return nil
}
