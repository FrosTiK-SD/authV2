package util

import (
	"github.com/FrosTiK-SD/auth/constants"
	"github.com/gin-gonic/gin"
)

func GetNoCache(ctx *gin.Context) bool {
	return ctx.GetHeader(constants.CACHE_CONTROL_HEADER) == constants.NO_CACHE
}
