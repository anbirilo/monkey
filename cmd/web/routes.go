package main

import (
	"net/http"
	//"snippetbox.anbirilo.net/ui"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	//fileServer := http.FileServer(http.FS(ui.Files))
	//router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	//router.HandlerFunc(http.MethodGet, "/ping", ping)

	dynamic := alice.New()

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	// router.Handler(http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.snippetView))
	// ... остальные маршруты остаются закомментированными ...

	// protected := dynamic.Append(app.requireAuthentication)
	// ... маршруты protected тоже остаются закомментированными ...

	standard := alice.New()

	return standard.Then(router)
}
