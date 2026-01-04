package service

import (
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

func (u *UserService) Create(newUser *types.User) error {
	return u.userRepository.Create(newUser)
}

func (u *UserService) Get() ([]*types.User, error) {
	return u.userRepository.Get()
}

func (u *UserService) Update(updatedUser *types.User) error {
	return u.userRepository.Update(updatedUser)
}

func (u *UserService) Delete(user *types.User) error {
	return u.userRepository.Delete(user.Name)
}
