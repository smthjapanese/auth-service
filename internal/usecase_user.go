package internal

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"register/internal/entity"
	"register/internal/repository"
)

type UserUseCase interface {
	Register(ctx context.Context, user entity.User) error
	Login(ctx context.Context, user entity.User) (string, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (u *userUseCase) Register(ctx context.Context, user entity.User) error {
	if len(user.Username) <= 3 || user.Username == "" {
		return entity.ErrInvalidUsername
	}
	if len(user.Password) < 6 || user.Password == "" {
		return entity.ErrInvalidPass
	}
	existingUser, err := u.repo.GetUserByUsername(ctx, user.Username)
	if err == nil && existingUser != nil {
		return entity.ErrUsernameAlreadyExist
	}
	if err != nil && !errors.Is(err, entity.ErrUsernameNotFound) {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	err = u.repo.CreateUser(ctx, &user)
	if err != nil {
		return err
	}
	return nil
}
func (u *userUseCase) Login(ctx context.Context, user entity.User) (string, error) {

}
