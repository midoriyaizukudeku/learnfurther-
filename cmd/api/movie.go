package main

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"time"

	"main/internal/data"
)

func (app *application) createmovie(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int      `json:"year"`
		Runtime int      `json:"runtime"`
		Genre   []string `json:"genre"`
	}

	err := app.Readjson(w, r, &input)
	if err != nil {
		app.badRequestresponse(w,r,err)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
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
		app.serverErrorResponse(w, r, err)
		return
	}

	//fmt.Fprintf(w, "RESULT FOR %d", id)
}
