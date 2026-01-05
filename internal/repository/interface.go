package repository

import (
	"context"

	"github.com/lunghyun/CRUD_SERVER/internal/types"
)

type UserRepository interface {
	Create(context.Context, *types.User) error
	Get(context.Context) ([]*types.User, error)
	Update(context.Context, *types.User) error
	Delete(context.Context, string) error
}
