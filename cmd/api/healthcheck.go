package main

import (
	//"encoding/json"
	//"fmt"
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"statius":     "availinble",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writejson(w, 200, data, nil)
	if err != nil {
		http.Error(w, "imo no", http.StatusInternalServerError)
		return
	}

	
}
