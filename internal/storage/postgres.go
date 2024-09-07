package storage

import (
	"context"
	"database/sql"

	"github.com/basado1991/jwt_auth_service/internal/types"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	conn *sql.DB
}

func NewPostgresStorage(addr string) (*PostgresStorage, error) {
	conn, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	return &PostgresStorage{conn: conn}, nil
}

func (s PostgresStorage) GetUserById(ctx context.Context, id string) (*types.User, error) {
	row := s.conn.QueryRowContext(ctx, "SELECT id, name, email, refresh_token FROM users WHERE id=$1", id)

	var u types.User

	if err := row.Scan(&u.Id, &u.Name, &u.Email, &u.RefreshToken); err != nil {
		return nil, err
	}

	return &u, nil
}

func (s PostgresStorage) UpdateUser(ctx context.Context, user *types.User) error {
	_, err := s.conn.ExecContext(ctx, "UPDATE users SET name=$1, email=$2, refresh_token=$3 WHERE id=$4", user.Name, user.Email, user.RefreshToken, user.Id)
	if err != nil {
		return err
	}

	return nil
}
