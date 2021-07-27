package tools

import (
	"github.com/gin-gonic/gin"
)

func Serialize(ctx *gin.Context) {
	body := make(map[string]interface{})
	err := ctx.BindJSON(&body)
	if err != nil {
		logger.Errorf("json serialized failed: %s", err.Error())
		ctx.JSON(200, err.Error())
	}
	logger.Info("json serialized successfully")
	ctx.JSON(200, body)
}
