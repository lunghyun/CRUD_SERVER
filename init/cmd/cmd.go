package cmd

import (
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

func NewCmd(filepath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filepath),
	}

	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)

	_ = c.network.ServerStart(c.config.Server.Port)

	return c
}
