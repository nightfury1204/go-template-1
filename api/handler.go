package api

import (
	"github.com/go-chi/chi"
)

func booksRouter(ctrl *BookController) chi.Router {
	h := chi.NewRouter()
	h.Group(func(r chi.Router) {
		r.Get("/", ctrl.ListBooks)
		r.Post("/", ctrl.CreateBook)
	})

	return h
}
