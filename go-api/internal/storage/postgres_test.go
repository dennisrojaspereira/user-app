package storage

import (
	"context"
	"createuserviper/go-api/internal/domain"
	"os"
	"testing"
)

func getTestPostgresStore(t *testing.T) *PostgresStore {
	conn := os.Getenv("POSTGRES_TEST")
	if conn == "" {
		t.Skip("POSTGRES_TEST not set")
	}
	store, err := NewPostgresStore(conn)
	if err != nil {
		t.Fatalf("failed to connect to postgres: %v", err)
	}
	return store
}

func TestPostgresStore_CreateUser(t *testing.T) {
	store := getTestPostgresStore(t)
	user, err := store.CreateUser(context.Background(), domain.CreateUserInput{Name: "Dave", Email: "dave@example.com"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.Name != "Dave" || user.Email != "dave@example.com" {
		t.Fatalf("unexpected user: %+v", user)
	}
}

func TestPostgresStore_GetUser(t *testing.T) {
	store := getTestPostgresStore(t)
	user, _ := store.CreateUser(context.Background(), domain.CreateUserInput{Name: "Eve", Email: "eve@example.com"})
	got, ok := store.GetUser(context.Background(), user.ID)
	if !ok {
		t.Fatal("user not found")
	}
	if got.Name != "Eve" {
		t.Fatalf("unexpected user: %+v", got)
	}
}

func TestPostgresStore_ListUsers(t *testing.T) {
	store := getTestPostgresStore(t)
	store.CreateUser(context.Background(), domain.CreateUserInput{Name: "Frank", Email: "frank@example.com"})
	users := store.ListUsers(context.Background())
	if len(users) == 0 {
		t.Fatalf("expected at least 1 user, got %d", len(users))
	}
}
