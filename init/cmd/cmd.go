package cmd

import (
	"fmt"

	"github.com/lunghyun/CRUD_SERVER/config"
)

type Cmd struct {
	config *config.Config
}

func NewCmd(filepath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filepath),
	}

	fmt.Println(c.config.Server.Port)

	return c
}
