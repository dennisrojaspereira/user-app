# CreateUserVIPER + Go API

## Structure
- Swift Package (VIPER) under `Sources/CreateUser` with XCTest under `Tests/CreateUserTests`
- Go API under `go-api` with unit tests and Dockerfile
- `docker-compose.yml` to run the API

## Run Go API
```bash
docker compose up --build
# API at http://localhost:8080
# POST /users  { "name": "...", "email": "..." }
# GET  /users
# GET  /users/{id}
# GET  /health
```

## Run Go tests
```bash
cd go-api
go test ./...
```

## Using from Swift
Point your networking to `http://localhost:8080/users` for create.
