package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        _, _ = w.Write([]byte(`{"status":"ok","service":"heran-music-api"}`))
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Heran-Music API listening on :%s", port)
    log.Fatal(http.ListenAndServe(":"+port, mux))
}
