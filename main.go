package main

import (
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"time"
)

const port = ":8080"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", DefaultEndpointHandler).Methods("GET")
	router.HandleFunc("/{id}", RedirectEndpointHandler).Methods("GET")
	router.HandleFunc("/create", CreationEndpointHandler).Methods("POST")
	http.ListenAndServe(port, router)
}
