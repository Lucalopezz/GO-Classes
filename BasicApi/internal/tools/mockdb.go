package tools

import "time"

// mockDB is an in-memory implementation of DatabaseInterface.
//
// It has no fields because the sample records are stored in package-level
// maps. A real implementation could replace this type with a repository that
// owns a SQL connection pool or an API client without changing the handlers.
type mockDB struct{}

// mockLoginDetails is the sample authentication table used by the mock
// database. The map key is the username, which makes lookups constant-time on
// average and keeps the example focused on the HTTP flow rather than storage.
var mockLoginDetails = map[string]LoginDetails{
	"alex":  {AuthToken: "token123", Username: "alex"},
	"jason": {AuthToken: "token456", Username: "jason"},
	"marie": {AuthToken: "token789", Username: "marie"},
}

// mockCoinDetails is the sample account-balance table used by the mock
// database.
var mockCoinDetails = map[string]CoinDetails{
	"alex":  {Coins: 100, Username: "alex"},
	"jason": {Coins: 200, Username: "jason"},
	"marie": {Coins: 300, Username: "marie"},
}

// GetUserLoginDetails retrieves a user's login record from the in-memory mock.
//
// The 100 ms delay simulates database latency so the request flow behaves more
// like an application that performs I/O. The map lookup returns a value copy;
// returning a pointer to that local copy is safe because Go keeps it alive as
// long as the caller uses the pointer.
func (db *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(100 * time.Millisecond) // Simulate database latency

	clientData, exists := mockLoginDetails[username]
	if !exists {
		return nil
	}

	return &clientData
}

// GetUserCoinDetails retrieves a user's coin record from the in-memory mock.
// It follows the same lookup and nil-on-missing-record convention as the login
// lookup above.
func (db *mockDB) GetUserCoinDetails(username string) *CoinDetails {
	time.Sleep(100 * time.Millisecond) // Simulate database latency

	clientData, exists := mockCoinDetails[username]
	if !exists {
		return nil
	}
	return &clientData
}

// SetupDatabase satisfies DatabaseInterface. The in-memory database needs no
// connection or migration step, so initialization succeeds immediately.
func (db *mockDB) SetupDatabase() error {
	return nil
}
