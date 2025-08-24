package storage

import (
	"context"
	"createuserviper/go-api/internal/domain"
	"testing"
)

func TestMemoryStore_CreateUser(t *testing.T) {
	store := NewMemoryStore()
	user, err := store.CreateUser(context.Background(), domain.CreateUserInput{Name: "Alice", Email: "alice@example.com"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.Name != "Alice" || user.Email != "alice@example.com" {
		t.Fatalf("unexpected user: %+v", user)
	}
}

func TestMemoryStore_GetUser(t *testing.T) {
	store := NewMemoryStore()
	user, _ := store.CreateUser(context.Background(), domain.CreateUserInput{Name: "Bob", Email: "bob@example.com"})
	got, ok := store.GetUser(context.Background(), user.ID)
	if !ok {
		t.Fatal("user not found")
	}
	if got.Name != "Bob" {
		t.Fatalf("unexpected user: %+v", got)
	}
}

func TestMemoryStore_ListUsers(t *testing.T) {
	store := NewMemoryStore()
	store.CreateUser(context.Background(), domain.CreateUserInput{Name: "Carol", Email: "carol@example.com"})
	users := store.ListUsers(context.Background())
	if len(users) != 1 {
		t.Fatalf("expected 1 user, got %d", len(users))
	}
}
