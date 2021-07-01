package main

import (
	"cainiaofundbackend/crontab"
	"cainiaofundbackend/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	logger.InitLogrus()
	crontab.InitCron()
	logrus.Info("dddds")
}
