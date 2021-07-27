package main

import (
	"fmt"
	"gin-demo/config"
	"gin-demo/router"
	_ "gin-demo/router/routerGroup"
)

func main() {
	err := router.Router.Run(config.PORT)
	if err != nil {
		fmt.Println(err)
	}
}
