package job

import (
	"cainiaofundbackend/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getCtx() *gin.Context {
	ctx := &gin.Context{}

	ctxID := uuid.NewString()
	logger.LogrusFormatter.SetCtxId(ctxID)

	ctx.Set("ctxID", ctxID)
	return ctx
}
