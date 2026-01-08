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

func (s *UserService) Create(ctx context.Context, newUser *types.User) error {
	return s.userRepository.Create(ctx, newUser)
	//tx, err := s.userRepository.
}

func (s *UserService) Get(ctx context.Context) ([]*types.User, error) {
	return s.userRepository.Get(ctx)
}

func (s *UserService) Update(ctx context.Context, updatedUser *types.User) error {
	return s.userRepository.Update(ctx, updatedUser)
}

func (s *UserService) Delete(ctx context.Context, user *types.User) error {
	return s.userRepository.Delete(ctx, user.Name)
}
