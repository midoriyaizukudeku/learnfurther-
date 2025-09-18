package main

import "net/http"

func (app *application) logError(r *http.Request , err error){
  var (
    method = r.Method
    url = r.URL.RequestURI()
    
  )
  app.logger.Error(err.Error(), "method", method, "url", url)
}

func(app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any){
  env := envolpe{"error":  message}

  err := app.writejson(w, status, env, nil)
  if err != nil {
    app.logError(r, err)
  }
}
