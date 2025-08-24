package storage

import (
    "context"
    "sync"
    "time"
    "createuserviper/go-api/internal/domain"
    "github.com/google/uuid"
)

type Store interface {
    CreateUser(ctx context.Context, in domain.CreateUserInput) (domain.User, error)
    GetUser(ctx context.Context, id string) (domain.User, bool)
    ListUsers(ctx context.Context) []domain.User
}

type MemoryStore struct {
    mu    sync.RWMutex
    users map[string]domain.User
}

func NewMemoryStore() *MemoryStore {
    return &MemoryStore{users: make(map[string]domain.User)}
}

func (s *MemoryStore) CreateUser(ctx context.Context, in domain.CreateUserInput) (domain.User, error) {
    id := uuid.NewString()
    u := domain.User{ID: id, Name: in.Name, Email: in.Email}
    s.mu.Lock()
    s.users[id] = u
    s.mu.Unlock()
    _ = ctx
    time.Sleep(10 * time.Millisecond)
    return u, nil
}

func (s *MemoryStore) GetUser(ctx context.Context, id string) (domain.User, bool) {
    s.mu.RLock()
    u, ok := s.users[id]
    s.mu.RUnlock()
    _ = ctx
    return u, ok
}

func (s *MemoryStore) ListUsers(ctx context.Context) []domain.User {
    s.mu.RLock()
    out := make([]domain.User, 0, len(s.users))
    for _, u := range s.users { out = append(out, u) }
    s.mu.RUnlock()
    _ = ctx
    return out
}
