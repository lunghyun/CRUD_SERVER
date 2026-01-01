package main

import (
	"flag"

	"github.com/lunghyun/CRUD_SERVER/init/cmd"
)

var configPathFlag = flag.String("config", "./env.toml", "config file path")

func main() {
	flag.Parse()
	cmd.NewCmd(*configPathFlag)
}
