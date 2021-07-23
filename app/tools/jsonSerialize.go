package tools

import (
	"github.com/gin-gonic/gin"
)

func Serialize(ctx *gin.Context) {
	body := make(map[string]interface{})
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(200, err.Error())
	}
	ctx.JSON(200, body)
}
