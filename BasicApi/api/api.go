// Package api contains the data structures and HTTP error helpers shared by
// the API layers of the application.
//
// Keeping these types in a small package gives handlers and middleware a
// common language for request parameters and JSON responses. This package does
// not register routes and does not know anything about the database; it only
// describes the HTTP-facing contract.
package api

import (
	"encoding/json"
	"net/http"
)

// CoinBalanceParams contains the values expected in the query string of the
// coin-balance endpoint.
//
// The current endpoint expects a request such as:
//
//	/accounts/coins?username=alex
//
// The authorization middleware reads the same username because it needs to
// identify which user's token should be checked before the handler runs.
type CoinBalanceParams struct {
	Username string
}

// CoinBalanceResponse is the successful JSON response returned by the
// coin-balance endpoint.
//
// Code is included in the body as part of this project's response convention,
// even though HTTP already carries the status code in the response line. The
// Balance field is an int64 so the API can represent balances larger than a
// platform-dependent int on every supported architecture.
type CoinBalanceResponse struct {
	// Code is the HTTP status code represented in the JSON body.
	Code int

	// Balance is the number of coins owned by the account.
	Balance int64
}

// Error is the common JSON shape used for unsuccessful requests.
type Error struct {
	// Code is the HTTP status code represented in the JSON body.
	Code int

	// Message is a human-readable description of the error.
	Message string
}

// writeError serializes an API error and writes the corresponding HTTP status
// code.
//
// net/http sends a successful status (200) automatically if a handler writes a
// body without calling WriteHeader first. Error responses must therefore call
// WriteHeader before encoding JSON. The Content-Type header tells clients that
// the body follows the JSON media type rather than plain text.
func writeError(w http.ResponseWriter, code int, message string) {
	response := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}

// RequestErrorHandler converts a client-side/request error into a 400 Bad
// Request response. It is a variable so applications or tests can replace the
// policy without changing every middleware or handler that uses it.
var RequestErrorHandler = func(w http.ResponseWriter, err error) {
	writeError(w, http.StatusBadRequest, err.Error())
}

// InternalErrorHandler converts an internal failure into a 500 response.
//
// The original error is intentionally not exposed to the client. Returning
// internal implementation details can leak database, filesystem, or service
// information. The error is still available to the caller for logging before
// this handler is invoked.
var InternalErrorHandler = func(w http.ResponseWriter, err error) {
	writeError(w, http.StatusInternalServerError, "An Unexpected error occurred")
}
