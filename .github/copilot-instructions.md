# Copilot Instructions for goprj001

## Project Overview
Energy and gas meter reading collection service - a clean web service that accepts JSON readings, stores them in-memory, and provides retrieval via REST API. Built with Go 1.25.5, standard library only (no external frameworks).

## Architecture & Data Flow

### Layered Design: DB Interface → Business Logic → Handlers → Server
- **Storage Layer** (`storage/IMDB.go`): In-memory database implementing interfaces for `AddGas()`, `AddEnergy()`, `ReadGas()`, `ReadEnergy()`
- **Application Layer** (`application/`): Use cases (`EnergyDataCase`, `GasDataCase`) that validate data per business rules and call storage
- **Handler Layer** (`handlers/`): HTTP handler functions that parse JSON, delegate to use cases, return HTTP responses
- **Server Layer** (`server/server.go`): Dependency injection container creating Logger, DB, UseCases; `router.go` maps routes

Example flow: HTTP Request → Handler → UseCase validation → Storage write → Response

### Critical Integration Point
[server/server.go](server/server.go#L26-L27): UseCases struct binds use cases to their storage implementations. This is where database swapping (e.g., PostgreSQL replacement for IMDB) would occur.

## Key Conventions

### Error Handling with Custom Types
Errors defined in [metainf/errors.go](metainf/errors.go): `ErrWrongData`, `ErrDBConn`, `ErrDBRead`. Use `errors.Is()` pattern to distinguish errors (see [handlers/sendenergy.go](handlers/sendenergy.go#L30-L38)).

### Generics Pattern
[utils/jsonparser.go](utils/jsonparser.go): Generic `ParseUserData[T any]()` function for type-safe JSON unmarshaling across handlers.

### Handler Constructor Pattern
Each handler requires a use case injected: `NewGasHandlerFunc(uc.GasUc) *GasHandler` returns struct with receiver methods (e.g., `ParseGasData()`, `GasHistory()`). This enables testability and loose coupling.

### Data Structures
- `DataGas`: `{value, data}` - single meter reading
- `DataEnergy`: `{day, night, summ, data}` - three-part reading where validation requires `day + night == summ`

### Business Validation
[application/applyenergy.go](application/applyenergy.go#L15-L17): Validation happens in use cases, not handlers. Energy readings must satisfy: `ed.Day >= 0 && ed.Night >= 0 && ed.Day + ed.Night == ed.Summ`.

## Routes & Endpoints
- `POST /sendgas` - Accept gas reading (JSON: `{value, data}`)
- `POST /sendenergy` - Accept energy reading (JSON: `{day, night, summ, data}`)
- `GET /history/gas` - Return all gas readings as JSON
- `GET /history/energy` - Return all energy readings as JSON

## Build & Deployment

### Docker Multi-Stage Build
[Dockerfile](Dockerfile): Builds with `CGO_ENABLED=0` and `scratch` base image for minimal binary (no libc dependency). Result: single executable in `/app/myapp.goprj001`.

### Local Development
```bash
go run main.go  # Server on :8081, logs to stdout via slog
```

### Docker Compose
[compose.yaml](compose.yaml): Backend on port 8081, UI (HTML frontend) on port 8082. Build via `docker compose up`.

## Development Status & Patterns
Project uses structured logging with `log/slog` (see [main.go](main.go#L14)). CORS middleware wraps entire router (not individual handlers). Recent refactor separated concerns: DB→Logic→Handler→Server. Next phases: add validation, logging layers, and data persistence.

## Important Notes
- **Language**: Russian comments throughout codebase
- **Middleware placement**: Always wrap router, never individual handlers (see [main.go](main.go#L18))
- **Time handling**: Use `time.Now()` for all timestamps in use cases (see [application/applyenergy.go](application/applyenergy.go#L18))
