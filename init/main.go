package main

import (
	"flag"

	"github.com/lunghyun/CRUD_SERVER/init/cmd"
)

var configPathFlag = flag.String("config", "../.env", "config file path")

func main() {
	flag.Parse()
	_, err := cmd.NewCmd(*configPathFlag)
	if err != nil {
		panic(err)
	}
}
