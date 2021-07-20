package crontab

import (
	"cainiaofundbackend/logger"
	"fmt"
	"github.com/sirupsen/logrus"
)

type CronLog struct {
	Logger *logrus.Logger
}

func NewCronLogger() CronLog {
	return CronLog{
		Logger: logger.NewLogger(),
	}
}

func (l CronLog) Info(msg string, keysAndValues ...interface{}) {
	m := fmt.Sprintf("cron:msg: %v", msg)
	str := kV2Str(keysAndValues...)

	l.Logger.Info(m + str)
}

func (l CronLog) Error(err error, msg string, keysAndValues ...interface{}) {
	m := fmt.Sprintf("cron:err:%v; msg: %v", err, msg)
	str := kV2Str(keysAndValues...)

	l.Logger.Error(m + str)
}

func kV2Str(keysAndValues ...interface{}) string {
	str := ""
	for k, v := range keysAndValues {
		if k/2 == 0 {
			str += fmt.Sprintf("%v:", v)
		} else {
			str += fmt.Sprintf("%v,", v)
		}
	}

	if str != "" {
		str = "; detail:" + str
	}

	return str
}
