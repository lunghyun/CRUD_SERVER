package repository

import (
	"database/sql"
	"fmt"

	"github.com/lunghyun/CRUD_SERVER/types"
)

type UserSqlRepository struct {
	db *sql.DB
}

func newUserSqlRepository(db *sql.DB) *UserSqlRepository {
	return &UserSqlRepository{
		db: db,
	}
}

func (u *UserSqlRepository) Create(newUser *types.User) error {
	//u.userMap = append(u.userMap, newUser)
	query := `INSERT INTO users (name, age) VALUES (?, ?)`
	_, err := u.db.Exec(query, newUser.Name, newUser.Age)
	if err != nil {
		return fmt.Errorf("user 생성 실패: %w", err)
	}

	return nil
}

func (u *UserSqlRepository) Get() []*types.User {
	// SELECT * FROM users
	query := `SELECT name, age FROM users`

	rows, err := u.db.Query(query)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var users []*types.User
	for rows.Next() {
		user := &types.User{}
		if err = rows.Scan(&user.Name, &user.Age); err != nil {
			continue
		}
		users = append(users, user)
	}

	return users
}

func (u *UserSqlRepository) Update(updatedUser *types.User) error {
	// name이 같은 user를 찾고, 수정
	query := `UPDATE users SET age = ? WHERE name = ?`

	_, err := u.db.Exec(query, updatedUser.Age, updatedUser.Name)
	if err != nil {
		return fmt.Errorf("user 수정 실패: %w", err)
	}

	return nil
}

func (u *UserSqlRepository) Delete(userName string) error {
	// name에 해당하는 user 삭제
	query := `DELETE FROM users WHERE name = ?`

	_, err := u.db.Exec(query, userName)
	if err != nil {
		return fmt.Errorf("user 삭제 실패: %w", err)
	}

	return nil
}
