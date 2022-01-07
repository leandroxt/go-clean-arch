package main

import (
	"errors"
	"log"
	"net/http"
)

func (app *application) serve() error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	log.Println("Starting server")

	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
