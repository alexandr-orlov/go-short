package routers

import (
	"github.com/alexandr-orlov/go-short/internal/handlers"
	"github.com/go-chi/chi"
)

func URLRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/", handlers.GetRootHandler)
	r.Get("/{id}", handlers.GetHandler)
	r.Post("/", handlers.PostHandler)

	return r
}
