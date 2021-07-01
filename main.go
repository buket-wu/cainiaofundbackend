package main

import (
	"cainiaofundbackend/crontab"
	"cainiaofundbackend/logger"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	logger.InitLogrus()
	crontab.InitCron()
	logrus.Info("dddds")

	time.Sleep(5 * time.Minute)
}
