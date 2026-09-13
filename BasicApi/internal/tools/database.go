// Package tools contains the persistence abstraction used by the API.
//
// The current implementation is intentionally small and in-memory so the
// example can run without an external database. The interface keeps the HTTP
// layers independent from that implementation.
package tools

import (
	log "github.com/sirupsen/logrus"
)

// LoginDetails is the authentication data associated with one user.
type LoginDetails struct {
	// AuthToken is the token accepted by the authorization middleware.
	AuthToken string
	// Username identifies the account that owns the token.
	Username string
}

// CoinDetails is the account data used by the coin-balance endpoint.
type CoinDetails struct {
	// Coins is the user's current coin balance.
	Coins int64
	// Username identifies the account that owns the balance.
	Username string
}

// DatabaseInterface defines the persistence operations required by the
// application.
//
// Interfaces describe behavior rather than storage. Both the authorization
// middleware and the coin-balance handler can work with any implementation
// that provides these methods, such as the current mock, a SQL repository, or
// a remote service client.
type DatabaseInterface interface {
	// GetUserLoginDetails looks up authentication data for a username. It
	// returns nil when the user does not exist.
	GetUserLoginDetails(username string) *LoginDetails
	// GetUserCoinDetails looks up the coin balance for a username. It returns
	// nil when the user does not exist.
	GetUserCoinDetails(username string) *CoinDetails
	// SetupDatabase initializes the underlying data source.
	SetupDatabase() error
}

// NewDatabase creates and initializes the application's database dependency.
//
// The return type is a pointer to an interface for compatibility with the
// existing callers. In idiomatic Go, returning DatabaseInterface directly is
// often simpler because interfaces already contain a reference to their
// concrete value. The current shape is documented here so callers understand
// why they dereference database before invoking an interface method.
func NewDatabase() (*DatabaseInterface, error) {
	// Assigning the concrete *mockDB to the interface means callers only see
	// the methods declared by DatabaseInterface.
	var database DatabaseInterface = &mockDB{}

	// SetupDatabase is part of the interface so a future implementation can
	// open connections, run migrations, or validate configuration here.
	var err error = database.SetupDatabase()
	if err != nil {
		log.Error("Error setting up database:", err)
		return nil, err
	}

	return &database, nil
}
