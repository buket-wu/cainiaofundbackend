package crontab

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func InitCron() {
	cronLog := NewCronLogger()
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	c := cron.New(
		cron.WithLogger(cronLog),
		cron.WithParser(secondParser),
	)

	id, err := c.AddJob("0/1 * * * * ?", GetSyncJob())
	if err != nil {
		logrus.Errorf("err:%v", err)
	}
	logrus.Info(id)

	c.Start()
}
