package main

import "net/http"

func(app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	w.Write([]byte("Saludos soy el primer endpoint, con informacion de esta aplicacion"))

}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about"{
		app.notFound(w)
		return
	}

	w.Write([]byte("En esta seccion se hablara un poco acerca de los antecedentes y el porque de esta aplicacion"))
}
