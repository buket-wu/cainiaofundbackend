package logger

import (
	"cainiaofundbackend/config"
	"fmt"
	"github.com/google/uuid"
	lF "github.com/jiajin1/logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var LogrusFormatter *lF.Formatter
var Logger *logrus.Logger

func init() {
	LogrusFormatter = lF.NewFormatter(false)
	LogrusFormatter.SetCtxId(uuid.NewString())

	logrus.SetLevel(getLevel())

	logrus.SetReportCaller(true)

	logrus.SetFormatter(LogrusFormatter)

	logrus.SetOutput(getOut())
}

func InitLogger() {
	Logger = NewLogger()
}

func NewLogger() *logrus.Logger {
	return &logrus.Logger{
		Level:        getLevel(),
		ReportCaller: true,
		Formatter:    LogrusFormatter,
		Out:          getOut(),
	}
}

func getLevel() logrus.Level {
	if config.Config.Trace {
		return logrus.TraceLevel
	}

	if config.Config.Debug {
		return logrus.DebugLevel
	}

	return logrus.InfoLevel
}

func getOut() io.Writer {
	if os.Getenv("CMD_OUT") == "1" {
		return os.Stdout
	}

	out, err := rotatelogs.New(
		config.Config.LogPath,
		rotatelogs.WithLinkName(config.Config.LogPath),
		rotatelogs.WithMaxAge(time.Duration(30*24)*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		panic(fmt.Sprintf("log init err:%v", err))
	}

	return out
}
