package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.GET("/api/giphy/:tag", app.handler.Giphy.GetGiphyByTag)

	return router
}
