package repository

import "github.com/lunghyun/CRUD_SERVER/types"

type UserRepository struct {
	UserMap []*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{}
}
