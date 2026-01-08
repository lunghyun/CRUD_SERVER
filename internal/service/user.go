package service

import (
	"context"

	"github.com/lunghyun/CRUD_SERVER/internal/repository"
	"github.com/lunghyun/CRUD_SERVER/internal/types"
)

type UserService struct {
	repo *repository.UserRepository
}

func newUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		repo: userRepository,
	}
}

func (s *UserService) Create(ctx context.Context, newUser *types.User) error {
	return s.repo.Create(ctx, newUser)
	//tx, err := s.repo.
}

func (s *UserService) Get(ctx context.Context) ([]*types.User, error) {
	return s.repo.Get(ctx)
}

func (s *UserService) Update(ctx context.Context, updatedUser *types.User) error {
	return s.repo.Update(ctx, updatedUser)
}

func (s *UserService) Delete(ctx context.Context, user *types.User) error {
	return s.repo.Delete(ctx, user.Name)
}
