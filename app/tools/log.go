package tools

import (
	"gin-demo/utils"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = utils.GetLogger("tools")
}
