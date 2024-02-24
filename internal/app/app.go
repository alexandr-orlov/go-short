package app

import (
	"net/http"

	"github.com/alexandr-orlov/go-short/internal/handlers"
	"github.com/go-chi/chi"
)

func UrlRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", handlers.GetRootHandler)
	r.Get("/{id}", handlers.GetHandler)
	r.Post("/", handlers.PostHandler)

	return r
}

func Run() {

	http.ListenAndServe(":8080", UrlRouter())

}
