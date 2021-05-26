package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.homedepot.com/EMC4JQ2/docker-go-api/products"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to products api.  Try /api/products or /api/products/id"))
}

func main() {
	PORT := ":" + os.Getenv("PORT")
	if PORT == ":" {
		PORT = "localhost:8080" 
	}
	// localhost:8080"
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	product := products.Product{}
	s := r.PathPrefix("/api/products").Subrouter()
	s.HandleFunc("", product.GetAll).Methods("GET")
	s.HandleFunc("", product.PostNew).Methods("POST")
	s.HandleFunc("/{id}", product.GetOne).Methods("GET")
	s.HandleFunc("/{id}", product.Update).Methods("PUT", "PATCH")
	s.HandleFunc("/{id}", product.Remove).Methods("DELETE")
	log.Fatal(http.ListenAndServe(PORT, r))
}
