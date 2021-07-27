package utils

import (
	"gin-demo/config"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

func GetLogger(name string) *logrus.Logger {
	logger := logrus.New()
	logPath := path.Join(config.Dir, "logs", name)
	CreatePath(logPath)
	file, err := os.OpenFile(path.Join(logPath, name+".log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	logger.SetOutput(file)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetReportCaller(true)

	return logger
}
