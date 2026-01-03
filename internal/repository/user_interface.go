package repository

import (
	"github.com/lunghyun/CRUD_SERVER/internal/types"
)

type UserRepository interface {
	Create(*types.User) error
	Get() []*types.User
	Update(*types.User) error
	Delete(string) error
}
