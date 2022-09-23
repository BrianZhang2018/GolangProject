package api

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func writeResponse(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Println(err)
	}
}

type AppError func(http.ResponseWriter, *http.Request) error

func (fn AppError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		pe, ok := err.(*MyError)
		if !ok {
			pe = &MyError{"1", "Server Error", 500}
			pe.DeveloperText = fmt.Sprintf("non standard API error: %s", err.Error())
		}
		writeResponse(w, pe.StatusCode, pe)
	}
}
