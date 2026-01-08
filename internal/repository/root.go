package repository

import (
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
	User *UserRepository
}

func NewRepository(db *sql.DB) *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			User: newUserRepository(db), // 구현체는 여기 주입
		}
	})

	return repositoryInstance
}
