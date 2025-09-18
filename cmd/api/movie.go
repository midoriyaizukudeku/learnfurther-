package main

import (
	"fmt"
	"net/http"
	"time"

	"main/internal/data"
)

func (app *application) createmovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "creating a movie...")
}

func (app *application) viewmovie(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		Id:        id,
		CreatedAt: time.Now(),
		Title:     "Vanilla sky",
		Runtime:   102,
		Genre:     []string{"suspense", "drama", "thrill"},
		Version:   1,
	}

	err = app.writejson(w, http.StatusOK, envolpe{"movie": movie}, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "sorry", http.StatusInternalServerError)
		return
	}

	//fmt.Fprintf(w, "RESULT FOR %d", id)
}
