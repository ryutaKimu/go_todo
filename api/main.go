package main

import (
	"go_todo_app/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	log.Println("Server starting on :8000")
	http.ListenAndServe(":8000", r)
}
