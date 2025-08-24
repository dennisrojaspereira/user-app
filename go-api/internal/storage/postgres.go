package storage

import (
	"context"
	"createuserviper/go-api/internal/domain"
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(connStr string) (*PostgresStore, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) CreateUser(ctx context.Context, in domain.CreateUserInput) (domain.User, error) {
	var id string
	err := s.db.QueryRowContext(ctx, `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`, in.Name, in.Email).Scan(&id)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{ID: id, Name: in.Name, Email: in.Email}, nil
}

func (s *PostgresStore) GetUser(ctx context.Context, id string) (domain.User, bool) {
	var u domain.User
	err := s.db.QueryRowContext(ctx, `SELECT id, name, email FROM users WHERE id = $1`, id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return domain.User{}, false
	}
	return u, true
}

func (s *PostgresStore) ListUsers(ctx context.Context) []domain.User {
	rows, err := s.db.QueryContext(ctx, `SELECT id, name, email FROM users`)
	if err != nil {
		return nil
	}
	defer rows.Close()
	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err == nil {
			users = append(users, u)
		}
	}
	return users
}
