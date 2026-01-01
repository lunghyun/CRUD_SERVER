package service

import "github.com/lunghyun/CRUD_SERVER/repository"

type UserService struct {
	userRepository *repository.UserRepository
}

func newUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}
