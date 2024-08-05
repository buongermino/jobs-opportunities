package main

import (
	"github.com/buongermino/jobs-opportunities.git/config"
	"github.com/buongermino/jobs-opportunities.git/router"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}

	router.Initialize()
}
