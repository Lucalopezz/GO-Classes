# Basic API

Basic API is a small Go HTTP service that demonstrates routing, middleware,
authentication, an API response contract, and a replaceable database layer.
The application currently uses an in-memory mock database, so it can run
without MySQL, PostgreSQL, Redis, or any other external service.

All examples below assume the project is run from the repository root.

## What the application does

The service exposes one protected endpoint:

```text
GET /accounts/coins?username=<username>
```

The request must include the user's mock token in the `Authorization` header.
For example:

```bash
curl -i \
  -H "Authorization: token123" \
  "http://localhost:8080/accounts/coins?username=alex"
```

Successful response:

```json
{
  "Code": 200,
  "Balance": 100
}
```

The current mock credentials are:

| Username | Authorization token | Balance |
| --- | --- | ---: |
| `alex` | `token123` | 100 |
| `jason` | `token456` | 200 |
| `marie` | `token789` | 300 |

These values are demonstration data only. They must not be treated as secure
credentials in a production system.

## Request lifecycle

An HTTP request travels through the application in this order:

1. `cmd/api/main.go` creates a `chi.Mux` router and starts the standard Go
   `net/http` server on `localhost:8080`.
2. `internal/handlers/api.go` registers the global `StripSlashes` middleware
   and the `/accounts` route group.
3. `internal/middleware/authorization.go` runs for every route inside the
   `/accounts` group. It reads `username` from the query string and the token
   from the `Authorization` header.
4. The middleware creates a `tools.DatabaseInterface` and compares the supplied
   token with the user's `AuthToken`.
5. If authentication fails, the middleware writes a `400 Bad Request` JSON
   error and stops the request.
6. If authentication succeeds, `internal/handlers/get_coin_balance.go` looks
   up the user's `CoinDetails` through the same database abstraction.
7. The handler maps the internal database record to the public
   `api.CoinBalanceResponse` and encodes it as JSON.

The important separation is that routing decides *where* a request goes,
middleware decides whether it may proceed, handlers implement the use case,
and the tools package supplies data-access behavior.

## Project structure

```text
.
├── api/
│   └── api.go
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── handlers/
│   │   ├── api.go
│   │   └── get_coin_balance.go
│   ├── middleware/
│   │   └── authorization.go
│   └── tools/
│       ├── database.go
│       └── mockdb.go
├── go.mod
├── go.sum
└── README.md
```

### `cmd/api`

This is the executable entry point. The `main` function assembles the router
and starts the server. Keeping the executable in `cmd` is a common Go layout:
it separates application startup from reusable internal packages.

### `api`

This package defines the data structures exchanged at the HTTP boundary:
query parameters, success responses, and error responses. It also owns the
shared error-handler functions used by middleware and handlers.

The package intentionally does not know how users or balances are stored.
That keeps the public API contract independent from persistence details.

### `internal/handlers`

This package connects URLs to use-case code. `api.go` registers routes, while
`get_coin_balance.go` implements the coin-balance endpoint.

The `internal` directory is a Go visibility boundary: packages outside this
module cannot import these application-internal packages directly. This helps
keep routing, middleware, and storage details private to the service.

### `internal/middleware`

Middleware wraps an HTTP handler and can run code before or after the wrapped
handler. The authorization middleware checks credentials before allowing the
coin-balance handler to run.

The middleware follows the standard `net/http` shape:

```go
func(next http.Handler) http.Handler
```

This makes it compatible with chi and with the standard library's HTTP
interfaces.

### `internal/tools`

This package defines `DatabaseInterface`, the persistence boundary used by the
rest of the application. `mockDB` implements that interface with maps.

The interface is useful even in a small example because handlers do not need
to know whether data comes from maps, SQL, a remote API, or another source.
Only the implementation behind the interface needs to change when storage
changes.

## External libraries

The project uses a small number of dependencies.

### `github.com/go-chi/chi`

Chi is a lightweight HTTP router for Go. It provides:

- `chi.NewRouter()` to create a router;
- `Route` to group endpoints under a common path prefix;
- `Use` to attach middleware to a router or route group;
- `Get` to register a handler for HTTP GET requests.

Chi integrates with the standard library because its router implements
`http.Handler`. That means `http.ListenAndServe` can use it directly without a
custom server abstraction.

### `github.com/go-chi/chi/middleware`

This package supplies reusable middleware maintained alongside chi. The
application uses `StripSlashes` to normalize paths with trailing slashes,
avoiding duplicate route behavior for paths such as `/accounts/coins` and
`/accounts/coins/`.

### `github.com/sirupsen/logrus`

Logrus is used for structured and leveled logging. The application uses its
`Error` function to record authentication, database, and server-startup
failures. `log.SetReportCaller(true)` adds the source location to log entries,
which is useful while tracing this educational project.

The service currently logs plain messages. A larger application could add
fields such as request IDs, usernames, or operation names, while taking care
not to log secrets such as authorization tokens.

### Go standard library packages

The application also relies on standard packages:

- `net/http` provides the server, request, response-writer, HTTP methods, and
  status constants;
- `encoding/json` serializes Go structs into JSON response bodies;
- `errors` creates the sentinel and contextual errors used by the application;
- `time` simulates storage latency in the mock database;
- `fmt` prints the startup message.

## Running the service

Start the API with:

```bash
go run ./cmd/api
```

The server listens on:

```text
http://localhost:8080
```

Try a valid request:

```bash
curl -i \
  -H "Authorization: token456" \
  "http://localhost:8080/accounts/coins?username=jason"
```

Try an invalid token:

```bash
curl -i \
  -H "Authorization: invalid" \
  "http://localhost:8080/accounts/coins?username=jason"
```

## Validation commands

Run all package tests:

```bash
go test ./...
```

Run the Go static analyzer:

```bash
go vet ./...
```

Format all Go files:

```bash
gofmt -w $(find . -name '*.go' -not -path './vendor/*')
```

There are currently no test files, so `go test ./...` primarily verifies that
all packages compile. `go vet ./...` checks for common correctness issues that
the compiler alone does not report.

## Design notes and next steps

This is an educational service rather than a production authentication
system. A production-ready version would normally add HTTPS, a deliberate
token scheme, secret management, persistent storage, request validation,
structured error semantics, tests, graceful shutdown, and configuration from
environment variables rather than hard-coded values.

One implementation detail worth knowing is that `NewDatabase` currently
returns a pointer to an interface. That is retained for compatibility with the
current code, but idiomatic Go commonly returns `DatabaseInterface` directly.
Changing that would simplify callers by removing the explicit dereference;
the interface-based architecture itself would remain the same.
