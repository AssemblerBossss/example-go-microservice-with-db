package httpapi

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/AssemblerBossss/example-go-microservice-with-db/internal/domains"
)

type dateTimeResponse struct {
	DateTime string `json:"datetime"`
}

func CreateDateTimeHandler(dtService *domains.DateTimeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := RequestIDFromContext(r.Context())
		log.Printf("CreateDateTimeHandler called with requestID=%s, with method=%s and with path= %s", requestID, r.Method, r.URL.Path)

		unixSeconds := dtService.CurrentUnixSeconds()
		t := time.Unix(unixSeconds, 0).UTC()
		iso8601 := t.Format(time.RFC3339)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(dateTimeResponse{DateTime: iso8601})
	}
}
