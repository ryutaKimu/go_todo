package controller

import (
	"encoding/json"
	"net/http"
)

type HelloController struct{}

type HelloResponse struct {
	Message string `json:"message"`
}

func NewHelloController() *HelloController {
	return &HelloController{}
}

func (c *HelloController) GetHello(w http.ResponseWriter, r *http.Request) {
	resp := HelloResponse{
		Message: "Hello, World!",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
