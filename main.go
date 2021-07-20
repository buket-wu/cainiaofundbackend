package main

import (
	"cainiaofundbackend/api"
	"cainiaofundbackend/config"
	"cainiaofundbackend/crontab"
)

func main() {
	crontab.InitCron()

	err := api.GetServer().Run(config.GetServerPort())
	if err != nil {
		panic(err)
	}
}
