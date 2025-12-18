package cmd

import (
	"fmt"

	"github.com/lunghyun/CRUD_SERVER/config"
	"github.com/lunghyun/CRUD_SERVER/network"
)

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd(filepath string) *Cmd {
	c := &Cmd{
		config:  config.NewConfig(filepath),
		network: network.NewNetwork(),
	}
	fmt.Println(c.config.Server.Port)

	_ = c.network.ServerStart(c.config.Server.Port)

	return c
}
