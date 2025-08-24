run:
	cd go-api && go run main.go

test:
	cd go-api && go test ./...

docker-build:
	cd go-api && docker build -f Dockerfile -t go-api .

docker-run:
	cd go-api && docker compose up --build

coverage-go:
	cd go-api && go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html

coverage-swift:
	cd Sources/CreateUser && swift test --enable-code-coverage && cd ../../ && xcrun llvm-cov show .build/debug/CreateUserPackageTests.xctest/Contents/MacOS/CreateUserPackageTests -instr-profile .build/debug/codecov/default.profdata > swift-coverage.txt
