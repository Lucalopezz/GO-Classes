package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Lucalopezz/Go-Classes/BasicApi/api"
	"github.com/Lucalopezz/Go-Classes/BasicApi/internal/tools"
	log "github.com/sirupsen/logrus"
)

// GetCoinBalance handles GET /accounts/coins requests.
//
// It reads the username from the query string, obtains the account data
// through DatabaseInterface, and maps the internal CoinDetails model to the
// public CoinBalanceResponse JSON contract.
func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	// Query parameters are exposed by net/http through URL.Query(). The
	// authorization middleware has already required a username, but the
	// handler reads it again because it needs the same value to query the
	// database for the account's balance.
	params := api.CoinBalanceParams{Username: r.URL.Query().Get("username")}
	var err error

	// NewDatabase returns a DatabaseInterface so the handler depends on the
	// abstraction rather than a concrete database implementation. The current
	// implementation is an in-memory mock, but the handler can keep the same
	// logic when a real database is introduced.
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w, err)
		return
	}
	// A map lookup returns a pointer to a CoinDetails value when the user
	// exists, or nil when no record is available. The nil check prevents a
	// panic and turns a missing record into an API error response.
	coinDetails := (*database).GetUserCoinDetails(params.Username)
	if coinDetails == nil {
		log.Error("Error getting coin details for user: ", params.Username)
		api.InternalErrorHandler(w, errors.New("coin details not found"))
		return
	}

	// The API response is deliberately different from the internal database
	// model. This prevents implementation details such as Username from being
	// exposed unless the public API contract explicitly requires them.
	response := api.CoinBalanceResponse{
		Balance: coinDetails.Coins,
		Code:    http.StatusOK,
	}

	// Encoding through json.Encoder writes a JSON representation directly to
	// the response stream. Setting Content-Type before encoding lets clients
	// choose the correct parser. The encoder also appends a newline, which is
	// valid JSON and makes command-line responses easier to read.
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error("Error encoding response: ", err)
		api.InternalErrorHandler(w, err)
		return
	}
}
