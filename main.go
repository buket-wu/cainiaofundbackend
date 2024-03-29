package main

import (
	"cainiaofundbackend/api"
	"cainiaofundbackend/config"
	"cainiaofundbackend/crontab"
	"cainiaofundbackend/logger"
)

func main() {
	logger.InitLogger()
	crontab.InitCron()

	err := api.GetServer().Run(config.GetServerPort())
	if err != nil {
		panic(err)
	}
}
