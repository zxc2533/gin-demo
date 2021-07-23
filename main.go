package main

import (
	"fmt"
	"gin-demo/router"
	_ "gin-demo/router/routerGroup"
)

func main() {
	err := router.Router.Run()
	if err != nil {
		fmt.Println(err)
	}
}
