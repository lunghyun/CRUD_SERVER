package repository

import (
	"context"
	"database/sql"
	"sync"
)

// 3tier 아키텍처
// 세션 정의
// DB 연결 설정
var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

type Repository struct {
	db   *sql.DB
	User *UserRepository
}

func NewRepository(db *sql.DB) *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			db:   db,
			User: newUserRepository(db), // 구현체는 여기 주입
		}
	})

	return repositoryInstance
}

func (r *Repository) WithinTx(ctx context.Context, fn func(txRepo *Repository) error) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	txRepo := &Repository{
		db:   r.db,
		User: r.User.WithTx(tx),
	}

	if err = fn(txRepo); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
