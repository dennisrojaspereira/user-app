FROM golang:1.22-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download || true
RUN go get github.com/google/uuid@v1.6.0
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /api ./

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /api /api
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/api"]
