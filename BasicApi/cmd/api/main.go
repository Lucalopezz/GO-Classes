// Package main is the executable entry point for the Basic API service.
package main

import (
	"fmt"
	"net/http"

	"github.com/Lucalopezz/Go-Classes/BasicApi/internal/handlers"
	// chi provides the router used to match incoming HTTP requests to handlers.
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

// main builds the HTTP router and starts the API server.
//
// net/http owns the server loop: it accepts TCP connections, parses HTTP
// requests, and invokes the router with an http.ResponseWriter and
// *http.Request. The router then selects the middleware and endpoint handler
// for each request.
func main() {
	// Include the source file and line number in log entries. This is useful
	// while learning the request flow because errors can be traced back to the
	// exact layer that produced them.
	log.SetReportCaller(true)

	// chi.NewRouter returns a router compatible with net/http's Handler
	// interface. That compatibility is why it can be passed directly to
	// http.ListenAndServe below.
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting server on :8080")

	// ListenAndServe blocks while the server is running. Passing "localhost"
	// binds the service to the local machine instead of all network interfaces.
	// The returned error is normally only reached when the server cannot start
	// or is stopped unexpectedly.
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Error("Error starting server:", err)
	}
}
