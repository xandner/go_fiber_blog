package main

import (
	"blog/application"
	"blog/config"
)

func main() {
	cfg,err:= config.NewConfig()
	if err != nil {
		panic(err)
	}

	application.Run(cfg)
}
