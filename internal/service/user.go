package service

import (
	"context"

	"github.com/lunghyun/CRUD_SERVER/internal/repository"
	"github.com/lunghyun/CRUD_SERVER/internal/types"
)

type UserService struct {
	repo *repository.Repository
}

func newUserService(repo *repository.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(ctx context.Context, newUser *types.User) error {
	return s.repo.WithinTx(ctx, func(txRepo *repository.Repository) error {
		return txRepo.User.Create(ctx, newUser)
	})
}

func (s *UserService) Get(ctx context.Context) ([]*types.User, error) {
	// 단순 조회면 tx 사용 x
	return s.repo.User.Get(ctx)
}

func (s *UserService) Update(ctx context.Context, updatedUser *types.User) error {
	return s.repo.WithinTx(ctx, func(txRepo *repository.Repository) error {
		return txRepo.User.Update(ctx, updatedUser)
	})
}

func (s *UserService) Delete(ctx context.Context, user *types.User) error {
	return s.repo.WithinTx(ctx, func(txRepo *repository.Repository) error {
		return txRepo.User.Delete(ctx, user.Name)
	})
}
