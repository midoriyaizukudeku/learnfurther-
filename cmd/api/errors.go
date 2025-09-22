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
  if err != nil{
    app.logError(r, err)
    w.WriteHeader(500)
  }
}

func(app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error){
  app.logError(r, err)

  message := "Server encountered problem sorry"
  app.errorResponse(w,r, 500, message)
  
}

func(app *application) notFound(w http.ResponseWriter, r *http.Request){
  message := "not found"
  app.errorResponse(w,r,http.StatusNotFound, message)
  
}

func(app *application)  methodNotAllowed(w http.ResponseWriter, r *http.Request){
  message := "This http method is not supported"
  app.errorResponse(w,r,http.StatusMethodNotAllowed, message)
  
}

func(app *application) badRequestresponse(w http.ResponseWriter, r *http.Request, err error){
  app.errorResponse(w,r,http.StatusBadRequest, err.Error())
}

func(app *application) Failedvalidationerror(w http.ResponseWriter, r *http.Request, error map[string]string){
   app.errorResponse(w,r, http.StatusUnprocessableEntity, error)
}