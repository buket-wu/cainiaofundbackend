package job

import (
	"cainiaofundbackend/logger"
	"fmt"
	"github.com/petermattis/goid"
	"github.com/sirupsen/logrus"
)

type SyncFund struct {
	Name string
}

func (j SyncFund) Run() {
	logrus.Info(j.Name)
}

func (j SyncFund) SetCtxId() {
	logger.LogrusFormatter.SetCtxId(fmt.Sprintf("%d", goid.Get()))
}
