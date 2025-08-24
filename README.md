# User App - Go API

## Estrutura do Projeto
- API Go em `go-api/` (com testes unitários, Dockerfile e docker-compose)

## Como subir a aplicação Go

### Usando Docker Compose
```bash
cd go-api
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

## Banco de dados Postgres
Ao subir o projeto, crie a tabela de usuários no banco Postgres:

```sql
CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	email TEXT NOT NULL
);
```

---

test:
docker-build:
docker-run:
# Makefile exemplo

```makefile
run:
	cd go-api && go run main.go


	cd go-api && go test ./...

docker-build:
	cd go-api && docker build -f Dockerfile -t go-api .


	cd go-api && docker compose up --build
```

