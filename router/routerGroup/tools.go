package routerGroup

import (
	"gin-demo/app/tools"
	"gin-demo/router"
)

func init() {
	group := router.Router.Group("/tools")
	{
		group.POST("/json_serialize", tools.Serialize)
	}
}
