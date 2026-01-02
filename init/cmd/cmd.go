package cmd

import (
	"fmt"

	"github.com/lunghyun/CRUD_SERVER/config"
	"github.com/lunghyun/CRUD_SERVER/network"
	"github.com/lunghyun/CRUD_SERVER/repository"
	"github.com/lunghyun/CRUD_SERVER/service"
)

type Cmd struct {
	config     *config.Config
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

	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)

	if err = c.network.ServerStart(c.config.Server.Port); err != nil {
		return nil, fmt.Errorf("서버가 시작되지 못함: %w", err)
	}

	return c, nil
}
