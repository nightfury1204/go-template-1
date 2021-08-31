package utils

import (
	"context"

	"github.com/go-chi/chi/middleware"
)

func GetTracingID(ctx context.Context) string {
	return middleware.GetReqID(ctx)
}
