package main

import "github.com/gorilla/mux"

func routes() *mux.Router {
	r := mux.NewRouter()

	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/server-health", checkHealth)

	return s
}
