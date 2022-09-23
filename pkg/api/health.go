package api

import (
	"fmt"
	"net/http"
)

type HealthResponse struct {
	Body interface{} `json:"body"`
}

func Health(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Hit Health endpoint!")
	writeResponse(w, http.StatusOK, &HealthResponse{Body: "Hello World"})
	return nil
}
