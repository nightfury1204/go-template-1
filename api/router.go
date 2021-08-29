package api

import (
	"go-template/api/middleware"
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Cors())

	r.NotFound(NotFoundHandler)
	r.MethodNotAllowed(MethodNotAllowed)

	r.Route("/", func(rt chi.Router) {
		// health status
		rt.Get("/_status", HealthStatus)
	})

	return r
}

// NotFoundHandler handles when no routes match
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

// MethodNotAllowed handles when no routes match
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func HealthStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	return
}
