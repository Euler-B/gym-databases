package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) Routes() http.Handler {

	r := chi.NewRouter()

	r.Get("/", app.home)
	r.Get("/about", app.about)

	return r

}
