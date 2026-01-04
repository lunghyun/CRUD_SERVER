package service

import (
	"context"

	"github.com/lunghyun/CRUD_SERVER/internal/repository"
	"github.com/lunghyun/CRUD_SERVER/internal/types"
)

type UserService struct {
	userRepository repository.UserRepository
}

func newUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) Create(ctx context.Context, newUser *types.User) error {
	return u.userRepository.Create(ctx, newUser)
}

func (u *UserService) Get(ctx context.Context) ([]*types.User, error) {
	return u.userRepository.Get(ctx)
}

func (u *UserService) Update(ctx context.Context, updatedUser *types.User) error {
	return u.userRepository.Update(ctx, updatedUser)
}

func (u *UserService) Delete(ctx context.Context, user *types.User) error {
	return u.userRepository.Delete(ctx, user.Name)
}
