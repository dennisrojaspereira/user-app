# User App - Go API

## Estrutura do Projeto
- API Go em `go-api/` (com testes unitários e Dockerfile)
- `docker-compose.yml` para orquestração

## Como subir a aplicação Go

### Usando Docker Compose
```bash
docker compose up --build
```
A API estará disponível em: http://localhost:8080

### Usando Makefile
```bash
make run
```

## Como rodar os testes Go
```bash
make test
```

## Endpoints principais
- `POST /users`  `{ "name": "...", "email": "..." }`
- `GET /users`
- `GET /users/{id}`
- `GET /health`

## Requisitos
- Go instalado
- Docker instalado

---

# Makefile exemplo

```makefile
run:
	cd go-api && go run main.go

test:
	cd go-api && go test ./...

docker-build:
	cd go-api && docker build -f Dockerfile -t go-api .

docker-run:
	docker compose up --build
```

