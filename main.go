package main

import (
	"cainiaofundbackend/api"
	"cainiaofundbackend/config"
	"cainiaofundbackend/crontab"
	"cainiaofundbackend/logger"
)

func main() {
	crontab.InitCron()
	logger.InitLogger()

	err := api.GetServer().Run(config.GetServerPort())
	if err != nil {
		panic(err)
	}
}
