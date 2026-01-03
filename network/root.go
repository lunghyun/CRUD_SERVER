package network

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lunghyun/CRUD_SERVER/service"
)

type Network struct {
	engine  *gin.Engine
	server  *http.Server
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
	// http.Server 생성 및 저장 -> gin이 stop이 없어서, 여기서 저장하고 구현해야함
	n.server = &http.Server{
		Addr:    ":" + port,
		Handler: n.engine.Handler(),
	}

	// ListenAndServe는 blocking 해줌
	// 정상 종료 시 http.ErrServerClosed 반환
	if err := n.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("비정상 종료: %w", err)
	}
	return nil
}

func (n *Network) ServerStop(ctx context.Context) error {
	if n.server == nil {
		return nil
	}
	// context : 타임아웃 시간 제어
	return n.server.Shutdown(ctx)
}
