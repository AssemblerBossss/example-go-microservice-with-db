package httpapi

import (
	"net/http"

	"github.com/AssemblerBossss/example-go-microservice-with-db/internal/domains"
)

func RequestIDMiddleware(genID domains.IDGenerator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := genID.NewID()
			w.Header().Set(requestIDHeader, requestID)
			next.ServeHTTP(w, r)
		})
	}
}
