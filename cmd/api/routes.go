package main

import (
  "net/http"

  "github.com/julienschmidt/httprouter"
  //"golang.org/x/net/route"
)

func(app *application) Routes() http.Handler{
  router := httprouter.New()
  router.NotFound = http.HandlerFunc(app.notFound)
  router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)

  router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
  router.HandlerFunc(http.MethodPost, "/v1/movies", app.createmovie)
  router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.viewmovie)
  return router
}
