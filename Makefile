run:
	cd go-api && go run main.go

test:
	cd go-api && go test ./...

docker-build:
	cd go-api && docker build -f Dockerfile -t go-api .

docker-run:
	docker compose up --build
