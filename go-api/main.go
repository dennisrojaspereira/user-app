package main

import (
	"context"
	h "createuserviper/go-api/internal/http"
	"createuserviper/go-api/internal/storage"
	"log"
	"net/http"
	"os"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	var store storage.Store
	pgConn := os.Getenv("POSTGRES")
	if pgConn != "" {
		pgStore, err := storage.NewPostgresStore(pgConn)
		if err != nil {
			log.Fatalf("Failed to connect to Postgres: %v", err)
		}
		store = pgStore
	} else {
		store = storage.NewMemoryStore()
	}

	// OpenTelemetry setup
	ctx := context.Background()
	exporter, err := otlptracehttp.New(ctx)
	if err != nil {
		log.Fatalf("Failed to create OTLP exporter: %v", err)
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			"service.name", "go-api",
			attribute.String("environment", "dev"),
		)),
	)
	defer func() { _ = tp.Shutdown(ctx) }()
	otel.SetTracerProvider(tp)

	srv := h.NewServer(store)
	handler := otelhttp.NewHandler(srv.Router(), "go-api")
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
