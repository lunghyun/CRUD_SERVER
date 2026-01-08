package service

import (
	"sync"

	"github.com/lunghyun/CRUD_SERVER/internal/repository"
)

// network와 repository의 다리 역할
// 일반적으로 net에서 API 스펙을 구성하고 받아오면
// API에 대한 응답을 repository에 넘겨주는 역할을 한다.

// 만약 따로 처리가 필요하다면 service에서 처리 후 repository에 넘겨준다.

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	UserService *UserService
}

func NewService(rep *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance.UserService = newUserService(rep)
	})

	return serviceInstance
}
