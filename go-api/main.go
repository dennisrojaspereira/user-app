package main

import (
    "log"
    "net/http"
    "os"
    h "createuserviper/go-api/internal/http"
    "createuserviper/go-api/internal/storage"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" { port = "8080" }
    store := storage.NewMemoryStore()
    srv := h.NewServer(store)
    log.Fatal(http.ListenAndServe(":"+port, srv.Router()))
}
