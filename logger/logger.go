package logger

import (
	"cainiaofundbackend/config"
	"fmt"
	lF "github.com/jiajin1/logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

var LogrusFormatter *lF.Formatter

func InitLogrus() {
	LogrusFormatter = lF.NewFormatter(false)

	logrus.SetLevel(logrus.TraceLevel)

	logrus.SetReportCaller(true)

	logrus.SetFormatter(LogrusFormatter)

	out, err := rotatelogs.New(
		config.LogPath,
		rotatelogs.WithLinkName(config.LogPath),
		rotatelogs.WithMaxAge(time.Duration(30*24)*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		fmt.Printf("err:%v", err)
	}

	logrus.SetOutput(out)
}

func NewLogger() *logrus.Logger {
	out, err := rotatelogs.New(
		config.LogPath,
		rotatelogs.WithLinkName(config.LogPath),
		rotatelogs.WithMaxAge(time.Duration(30*24)*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		fmt.Printf("err:%v", err)
	}

	return &logrus.Logger{
		Level:        getLevel(),
		ReportCaller: true,
		Formatter:    LogrusFormatter,
		Out:          out,
	}
}

func getLevel() logrus.Level {
	if config.Trace {
		return logrus.TraceLevel
	}

	if config.Debug {
		return logrus.DebugLevel
	}

	return logrus.InfoLevel
}
