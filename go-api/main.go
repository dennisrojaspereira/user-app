package main

import (
	h "createuserviper/go-api/internal/http"
	"createuserviper/go-api/internal/storage"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	var store storage.Store
	pgConn := os.Getenv("POSTGRES")
	if pgConn != "" {
		pgStore, err := storage.NewPostgresStore(pgConn)
		if err != nil {
			log.Fatalf("Failed to connect to Postgres: %v", err)
		}
		store = pgStore
	} else {
		store = storage.NewMemoryStore()
	}
	srv := h.NewServer(store)
	log.Fatal(http.ListenAndServe(":"+port, srv.Router()))
}
