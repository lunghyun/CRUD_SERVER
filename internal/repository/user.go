package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lunghyun/CRUD_SERVER/db/sqlc"
	"github.com/lunghyun/CRUD_SERVER/internal/types"
	"github.com/lunghyun/CRUD_SERVER/internal/types/cerrors"
)

type UserRepository struct {
	queries *sqlc.Queries
}

func newUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		queries: sqlc.New(db),
	}
}

func (r *UserRepository) WithTx(tx *sql.Tx) *UserRepository {
	return &UserRepository{
		queries: r.queries.WithTx(tx),
	}
}

func (r *UserRepository) Create(ctx context.Context, newUser *types.User) error {
	if err := r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name: newUser.Name,
		Age:  int32(newUser.Age),
	}); err != nil {
		return fmt.Errorf("user 생성 실패: %w", err)
	}

	return nil
}

func (r *UserRepository) Get(ctx context.Context) ([]*types.User, error) {
	users, err := r.queries.GetAllUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("user 조회 실패: %w", err)
	}

	result := make([]*types.User, 0, len(users))
	for _, user := range users {
		result = append(result, &types.User{
			Name: user.Name,
			Age:  int(user.Age),
		})
	}

	return result, nil
}

func (r *UserRepository) Update(ctx context.Context, updatedUser *types.User) error {
	result, err := r.queries.UpdateUserAge(ctx, sqlc.UpdateUserAgeParams{
		Name: updatedUser.Name,
		Age:  int32(updatedUser.Age),
	})
	if err != nil {
		return fmt.Errorf("user 수정 실패: %w", err)
	}

	rows, err := result.RowsAffected()
	// 인프라 오류
	if err != nil {
		return fmt.Errorf("row affected 실패: %w", err)
	}

	// name에 해당하는 row가 없으면 -> err가 nil이라서 래핑 안함
	if rows == 0 {
		return cerrors.Errorf(cerrors.NotFoundUser, nil)
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, userName string) error {
	result, err := r.queries.DeleteUserByName(ctx, userName)
	if err != nil {
		return fmt.Errorf("user 삭제 실패: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("row affected 실패: %w", err)
	}

	if rows == 0 {
		return cerrors.Errorf(cerrors.NotFoundUser, nil)
	}

	return nil
}
