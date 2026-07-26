package main

import (
	"net/http"

	"github.com/AssemblerBossss/example-go-microservice-with-db/internal/domains"
	"github.com/AssemblerBossss/example-go-microservice-with-db/internal/infra"
	"github.com/AssemblerBossss/example-go-microservice-with-db/internal/protocol/httpapi"
)

func buildHttpHandler() http.Handler {
	idGen := infra.UUIDGenerator{}
	mux := http.NewServeMux()

	clock := infra.SystemClock{}
	dfSystemClockService := domains.NewDateTimeService(clock)

	mux.Handle("GET /", httpapi.CreateDateTimeHandler(dfSystemClockService))

	var handler http.Handler = mux
	handler = httpapi.RequestIDMiddleware(idGen)(handler)
	return handler

}
