package network

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/lunghyun/CRUD_SERVER/service"
)

type Network struct {
	engine *gin.Engine

	service *service.Service
}

func NewNetwork(service *service.Service) *Network {
	r := &Network{
		engine:  gin.New(),
		service: service,
	}
	newUserRouter(r, service.UserService) // gin.engine에 엔드포인트 등록
	return r
}

func (n *Network) ServerStart(port string) error {
	return n.engine.Run(":" + port)
}

func (n *Network) ServerStop(ctx context.Context) error {
	// TODO gin.Engine의 stop 메서드가 없으니 고민
	return nil
}
