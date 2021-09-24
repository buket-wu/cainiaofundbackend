package routine

import (
	"cainiaofundbackend/config"
	"github.com/panjf2000/ants/v2"
	"github.com/sirupsen/logrus"
)

var Pool *ants.Pool

func init() {
	var err error
	Pool, err = ants.NewPool(config.Config.RoutineNum)
	if err != nil {
		logrus.Fatalf("init rountine pool fail; err:%v", err)
	}
}
