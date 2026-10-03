package main

import (
	"net/http"
	// "github.com/julienschmidt/httprouter"
	// "snippetbox.anbirilo.net/internal/models"
	// "snippetbox.anbirilo.net/internal/validator"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Monkey"))
}
