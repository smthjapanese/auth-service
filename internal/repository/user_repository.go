package repository

import (
	"context"
	"register/internal/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
}
