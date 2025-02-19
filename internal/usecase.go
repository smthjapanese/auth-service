package internal

import (
	"context"
	"register/internal/entity"
)

type UserUseCase interface {
	Register(ctx context.Context, user entity.User) error
	Login(ctx context.Context, user entity.User) (string, error)
}
