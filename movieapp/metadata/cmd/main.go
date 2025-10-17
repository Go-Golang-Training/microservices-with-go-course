package main

import (
	"log"
	"net/http"

	"movieexample.com/metadata/internal/controller/metadata"
	httphandler "movieexample.com/metadata/internal/handler/http"
	"movieexample.com/metadata/internal/repository/memory"
)

func main() {
	log.Println("Starting the movie metadata service")

	// Dependency injection and composition root
	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httphandler.New(ctrl)

	// Register HTTP handler
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))

	// Start server with error handling
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
