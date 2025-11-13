package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	//"golang.org/x/net/route"
)

func (app *application) Routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFound)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.listMovieshandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createmovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.viewmovie)
	router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.Updatemovie)
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.DeletemovieHandler)
	return router
}
