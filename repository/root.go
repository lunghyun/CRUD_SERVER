package repository

import (
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
	User *UserMemRepository
}

func NewRepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			User: newUserMemRepository(),
		}
	})

	return repositoryInstance
}
