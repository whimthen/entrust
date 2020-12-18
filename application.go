package main

import (
	"entrust/router"
)

func main() {
	// TODO load project config from config center
	router.InitializeRouter()
	err := router.Engine.Run(":8080")
	if err != nil {
		println(err)
	}
}
