package main

import (
	"github.com/ryutaKimu/go_todo/api/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	log.Println("Server starting on :8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
