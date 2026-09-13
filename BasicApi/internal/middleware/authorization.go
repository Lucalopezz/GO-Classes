// Package middleware contains HTTP middleware shared by multiple routes.
package middleware

import (
	"errors"
	"net/http"

	"github.com/Lucalopezz/Go-Classes/BasicApi/api"
	"github.com/Lucalopezz/Go-Classes/BasicApi/internal/tools"
	log "github.com/sirupsen/logrus"
)

// UnAuthorizedError is returned when a request does not contain acceptable
// authentication credentials.
//
// It is a package variable so the same error value can be logged and sent to
// the request error handler without creating a new error for every branch.
var UnAuthorizedError = errors.New("unauthorized access")

// Authorization is an HTTP middleware that protects a downstream handler with
// a username and token check.
//
// Middleware follows the decorator pattern used by net/http: it receives the
// next handler and returns another handler. The returned handler can inspect a
// request before deciding whether next.ServeHTTP should be called. If the
// credentials are invalid, the request ends in this middleware and the
// protected endpoint is never executed.
//
// This example uses a plain Authorization header because the mock database
// stores plain sample tokens. Production authentication should use a deliberate
// token format and secure transport such as HTTPS; real tokens should also be
// compared and stored according to the application's security requirements.
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// The username identifies the account being accessed. The token comes
		// from the standard HTTP Authorization header. net/http exposes both
		// values without requiring a framework-specific request type.
		var username string = r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")
		var err error

		// Reject incomplete credentials before creating a database connection.
		// Returning here is important: calling next after writing an error could
		// produce multiple responses for one request.
		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		// The middleware depends on DatabaseInterface instead of directly
		// depending on mockDB. This is dependency inversion: authentication
		// logic asks for the operations it needs and does not know how records
		// are stored.
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w, err)
			return
		}

		loginDetails := (*database).GetUserLoginDetails(username)

		// A missing user and a mismatched token are intentionally handled the
		// same way. Revealing which part failed could help an attacker discover
		// valid usernames.
		if loginDetails == nil || loginDetails.AuthToken != token {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}
		// Only authenticated requests reach the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}
