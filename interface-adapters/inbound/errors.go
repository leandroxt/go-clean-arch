package inbound

import (
	"log"
	"net/http"
)

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}

	err := writeJSON(w, status, env, nil)
	if err != nil {
		log.Println("ERROR | ", err.Error())
		w.WriteHeader(500)
	}
}

func serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Println("ERROR | ", err.Error())

	message := "the server encountered a problem and could not process your request"
	errorResponse(w, r, http.StatusInternalServerError, message)
}
