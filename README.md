# User App - Go API

## Project Structure
- Go API in `go-api/` (with unit tests, Dockerfile, and docker-compose)

## How to start the Go application

### Using Docker Compose
```bash
cd go-api
docker compose up --build
```
The API will be available at: http://localhost:8080

### Using Makefile
```bash
make run
```

## How to run Go tests
```bash
make test
```

## Main Endpoints
- `POST /users`  `{ "name": "...", "email": "..." }`
- `GET /users`
- `GET /users/{id}`
- `GET /health`

## Requirements
- Go installed
- Docker installed

## Postgres Database
When starting the project, create the users table in Postgres:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL
);
```

---

# Makefile example

```makefile
run:
	cd go-api && go run main.go

test:
	cd go-api && go test ./...

docker-build:
	cd go-api && docker build -f Dockerfile -t go-api .

docker-run:
	cd go-api && docker compose up --build
```

