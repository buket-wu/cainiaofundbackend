package job

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type SyncFund struct {
	Name string
}

func (j SyncFund) Run() {
	logrus.Info(j.Name)
	fmt.Println(j.Name)
}
