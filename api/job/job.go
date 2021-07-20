package job

import (
	j "cainiaofundbackend/crontab/job"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Sync(ctx *gin.Context) {
	syncJob := j.SyncFund{}
	syncJob.Run()

	ctx.JSON(http.StatusOK, "sync end")
	return
}
