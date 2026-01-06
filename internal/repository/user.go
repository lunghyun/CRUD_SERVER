package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lunghyun/CRUD_SERVER/internal/types"
	"github.com/lunghyun/CRUD_SERVER/internal/types/cerrors"
)

type UserSqlRepository struct {
	db *sql.DB
}

func newUserSqlRepository(db *sql.DB) *UserSqlRepository {
	return &UserSqlRepository{
		db: db,
	}
}

func (u *UserSqlRepository) Create(ctx context.Context, newUser *types.User) error {
	//u.userMap = append(u.userMap, newUser)
	query := `INSERT INTO users (name, age) VALUES (?, ?)`
	_, err := u.db.ExecContext(ctx, query, newUser.Name, newUser.Age)
	if err != nil {
		return fmt.Errorf("user 생성 실패: %w", err)
	}

	return nil
}

func (u *UserSqlRepository) Get(ctx context.Context) ([]*types.User, error) {
	// SELECT * FROM users
	query := `SELECT id, name, age FROM users`

	rows, err := u.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("user 조회 실패: %w", err)
	}
	defer func() {
		_ = rows.Close()
	}()

	var users []*types.User
	for rows.Next() {
		user := &types.User{}
		if err = rows.Scan(&user.Name, &user.Age); err != nil {
			continue
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Next() 오류: %w", err)
	}

	return users, nil
}

func (u *UserSqlRepository) Update(ctx context.Context, updatedUser *types.User) error {
	// name이 같은 user를 찾고, 수정
	query := `UPDATE users SET age = ? WHERE name = ?`

	result, err := u.db.ExecContext(ctx, query, updatedUser.Age, updatedUser.Name)
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

func (u *UserSqlRepository) Delete(ctx context.Context, userName string) error {
	// name에 해당하는 user 삭제
	query := `DELETE FROM users WHERE name = ?`

	result, err := u.db.ExecContext(ctx, query, userName)
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
