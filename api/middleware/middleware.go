package middleware

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
)

// Middleware represents http handler middleware
type Middleware func(http.Handler) http.Handler

var RequestID = middleware.RequestID
