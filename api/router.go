package api

import (
	"errors"
	"fmt"
	"go-template/api/middleware"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
)

func NewRouter(bookCtrl *BookController) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Cors())

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(chimiddleware.Timeout(60 * time.Second))

	r.NotFound(NotFoundHandler)
	r.MethodNotAllowed(MethodNotAllowed)

	r.Route("/", func(rt chi.Router) {
		// health status
		rt.Get("/status", HealthStatus)

		rt.Route("/api/v1", func(h chi.Router) {
			h.Mount("/books", booksRouter(bookCtrl))
		})
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
}

func parseSkipLimit(r *http.Request, defaultLimit, max int) (int, int, int, error) {
	q := r.URL.Query()
	var page, skip, limit int = 1, 0, defaultLimit
	pageQ := q.Get("page")
	if pageQ != "" {
		pageInt, err := strconv.Atoi(pageQ)
		if err != nil {
			return page, skip, limit, errors.New("failed to parse page")
		}
		page = pageInt
	}

	limitQ := q.Get("limit")
	if limitQ != "" {
		limitInt, err := strconv.Atoi(limitQ)
		if err != nil {
			return page, skip, limit, errors.New("failed to parse limit")
		}
		limit = limitInt
	}

	if limit > max {
		limit = max
	}
	skip = limit * (page - 1)
	if skip < 0 {
		skip = 0
	}
	return page, skip, limit, nil
}

func getNextPreviousPagerLink(path string, page, limit, gotNumsItems int) (*string, *string) {
	var previous, next string
	if page-1 > 0 {
		previous = fmt.Sprintf("%s?limit=%d&page=%d", path, limit, page-1)
	}
	if gotNumsItems > limit {
		next = fmt.Sprintf("%s?limit=%d&page=%d", path, limit, page+1)
	}
	return &previous, &next
}
