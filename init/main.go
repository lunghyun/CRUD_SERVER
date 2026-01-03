package main

import (
	"flag"

	"github.com/lunghyun/CRUD_SERVER/init/cmd"
)

var configPathFlag = flag.String("config", "../.env", "config file path")

func main() {
	flag.Parse()
	c, err := cmd.NewCmd(*configPathFlag)
	if err != nil {
		panic(err)
	}

	// blocking + 내부 graceful shutdown
	if err = c.Run(); err != nil {
		panic(err)
	}
}
