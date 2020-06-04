package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	db "github.homedepot.com/EMC4JQ2/docker-go-api/database"
	"github.homedepot.com/EMC4JQ2/docker-go-api/products"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to products api.  Try /api/products or /api/products/id"))
}

func main() {
	conn := map[string]string{
		"host":     os.Getenv("DB_HOST"),
		"port":     os.Getenv("DB_PORT"),
		"user":     os.Getenv("DB_USER"),
		"pw":       os.Getenv("DB_PASSWORD"),
		"database": os.Getenv("DB_DATABASE"),
	}
	caller := db.DBCalls{}
	db.Init(caller, conn)

	product := products.GetHandlers(caller)

	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)

	s := r.PathPrefix("/api/products").Subrouter()
	s.HandleFunc("", product.GetAll).Methods("GET")
	s.HandleFunc("", product.PostNew).Methods("POST")
	s.HandleFunc("/{id}", product.GetOne).Methods("GET")
	s.HandleFunc("/{id}", product.Update).Methods("PUT", "PATCH")
	s.HandleFunc("/{id}", product.Remove).Methods("DELETE")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
