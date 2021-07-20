package crontab

import (
	"cainiaofundbackend/crontab/job"
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

	c.Start()

	addJob(c, "syncJob", "CRON_TZ=Asia/Shanghai 0 30 * * *", &job.SyncFund{})

}

func addJob(c *cron.Cron, name, spec string, cmd cron.Job) {
	id, err := c.AddJob(spec, cmd)
	if err != nil {
		logrus.Errorf("add %s job err:%v", name, err)
	}
	logrus.Infof("add %s success:id:%v", name, id)
}
