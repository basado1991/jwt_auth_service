package storage

import (
	"context"

	"github.com/basado1991/jwt_auth_service/internal/types"
)

type Storage interface {
  GetUserById(ctx context.Context, id string) (*types.User, error)
  UpdateUser(ctx context.Context, user *types.User) error
}
