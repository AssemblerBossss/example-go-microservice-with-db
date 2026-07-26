package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: buildHttpHandler(),
	}
	log.Printf("Starting server on port %s", srv.Addr)
	srv.ListenAndServe()
}
