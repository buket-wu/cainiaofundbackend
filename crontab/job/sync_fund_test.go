package job

import (
	"cainiaofundbackend/logger"
	"testing"
)

func init() {
	logger.NewLogger()
}

func TestSyncFund_Run(t *testing.T) {
	sync := SyncFund{}
	sync.Run()
}
