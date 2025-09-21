package main

import (
	//"encoding/json"
	//"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	env := envolpe{
		"availible": "yes",
		"system_info": map[string]string{
			"envi": app.config.env,
			"version": version,
		},
	}

	err := app.writejson(w, 200, env, nil)
	if err != nil {
		app.serverErrorResponse(w,r, err)
	}

	
}
