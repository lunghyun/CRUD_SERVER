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

// TODO tx구현 계획(db connection을 어떻게 주입할까/repo tx 메서드를 어떻게 구현할까)
//func (u *UserService) Create(ctx context.Context, db *sql.DB, newUser *types.User) error {
//	tx, err := db.BeginTx(ctx, nil)
//	if err != nil {
//		return err
//	}
//	defer tx.Rollback()
//	if err = u.userRepository.Create(ctx, tx, newUser); err != nil {
//		return err
//	}
//	return tx.Commit()
//}
//
//func (u *UserService) Get(ctx context.Context, db *sql.DB) ([]*types.User, error) {
//	tx, err := db.BeginTx(ctx, nil)
//	if err != nil {
//		return nil, err
//	}
//	defer tx.Rollback()
//	users, err := u.userRepository.Get(ctx, tx)
//	if err != nil {
//		return nil, err
//	}
//	return users, tx.Commit()
//}
//
//func (u *UserService) Update(ctx context.Context, db *sql.DB, updatedUser *types.User) error {
//	tx, err := db.BeginTx(ctx, nil)
//	if err != nil {
//		return err
//	}
//	defer tx.Rollback()
//	if err = u.userRepository.Update(ctx, tx, updatedUser); err != nil {
//		return err
//	}
//	return tx.Commit()
//}
//
//func (u *UserService) Delete(ctx context.Context, db *sql.DB, user *types.User) error {
//	tx, err := db.BeginTx(ctx, nil)
//	if err != nil {
//		return err
//	}
//	defer tx.Rollback()
//	if err = u.userRepository.Delete(ctx, tx, user.Name); err != nil {
//		return err
//	}
//	return tx.Commit()
//}
//
