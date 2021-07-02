package crontab

import (
	"github.com/robfig/cron/v3"
)

func InitCron() {
	cronLog := NewCronLogger()
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	c := cron.New(
		cron.WithLogger(cronLog),
		cron.WithParser(secondParser),
	)

	c.Start()
}
