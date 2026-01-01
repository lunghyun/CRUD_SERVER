package repository

import "github.com/lunghyun/CRUD_SERVER/types"

type UserRepository struct {
	userMap []*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{}
}
