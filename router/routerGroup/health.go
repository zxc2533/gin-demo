package routerGroup

import (
	"gin-demo/config"
	"gin-demo/router"

	"github.com/gin-gonic/gin"
)

func init() {
	router.Router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": config.Dir})
	})
}
