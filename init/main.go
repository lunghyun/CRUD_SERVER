package main

import (
	"flag"

	"github.com/lunghyun/CRUD_SERVER/init/cmd"
)

var configPathFlag = flag.String("config", "./config.toml", "config file not found")

func main() {
	flag.Parse()
	cmd.NewCmd(*configPathFlag)
}
