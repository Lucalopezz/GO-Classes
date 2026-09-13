// Package handlers contains the HTTP routing configuration and endpoint
// implementations for the application.
package handlers

import (
	"github.com/Lucalopezz/Go-Classes/BasicApi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
)

// Handler configures the application's routes on a chi router.
//
// A chi.Mux is a router that implements net/http.Handler. It is responsible
// for matching the request method and path, while middleware adds behavior
// around the selected endpoint. The order here matters: middleware registered
// on a route runs before that route's handler.
func Handler(r *chi.Mux) {
	// StripSlashes normalizes paths by removing a trailing slash. For example,
	// /accounts/coins/ is treated like /accounts/coins. This is provided by
	// chi's middleware package rather than implemented manually in this app.
	r.Use(chimiddleware.StripSlashes)

	// Route groups let related endpoints share a path prefix and middleware.
	// Every endpoint inside /accounts is protected by Authorization before its
	// own handler is called.
	r.Route("/accounts", func(router chi.Router) {
		router.Use(middleware.Authorization)

		// GET is used because this endpoint reads account data without changing
		// the account. The username is supplied as a query parameter.
		router.Get("/coins", GetCoinBalance)
	})
}
