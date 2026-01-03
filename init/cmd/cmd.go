package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lunghyun/CRUD_SERVER/internal/config"
	"github.com/lunghyun/CRUD_SERVER/internal/infra"
	"github.com/lunghyun/CRUD_SERVER/internal/network"
	"github.com/lunghyun/CRUD_SERVER/internal/repository"
	"github.com/lunghyun/CRUD_SERVER/internal/service"
)

type Cmd struct {
	config     *config.Config
	database   *infra.DB
	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

func NewCmd(filepath string) (*Cmd, error) {
	c := &Cmd{}

	cfg, err := config.NewConfig(filepath)
	if err != nil {
		return nil, fmt.Errorf("config 불러오기 실패: %w", err)
	}
	c.config = cfg

	dbConn, err := infra.NewDB(c.config.Database)
	if err != nil {
		return nil, fmt.Errorf("DB 연결 실패: %w", err)
	}
	c.database = dbConn

	c.repository = repository.NewRepository(c.database.Conn)
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)

	return c, nil
}

func (c *Cmd) Run() error {
	// 고루틴으로 변경 -> blocking에 의지되는 상태 해제

	// 서버 에러를 받는 채널
	errChan := make(chan error, 1)

	// 1. 서버 시작
	go func() {
		if err := c.network.ServerStart(c.config.Server.Port); err != nil {
			errChan <- err
		}
	}()

	// 2. signal 대기 : ctrl+c, kill
	quitSign := make(chan os.Signal, 1)
	signal.Notify(quitSign, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		log.Printf("서버 시작 실패: %v\n", err)
		return err
	case <-quitSign:
		log.Println("서버 종료 중...")
	}

	// 3. Graceful shutdown (timeout: 5s)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := c.network.ServerStop(ctx); err != nil {
		log.Printf("서버 종료 에러: %v\n", err)
	}

	// 4. DB shutdown
	if err := c.database.Close(); err != nil {
		log.Printf("DB close 에러: %v\n", err)
	}

	return nil
}
