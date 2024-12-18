package main

import (
	"log"
	"net/http"

	"github.com/yusufhalabi/quran-search/internal/api/router"
)

func main() {
	r := router.Setup()

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
