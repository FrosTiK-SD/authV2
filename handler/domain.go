package handler

import (
	"net/http"

	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/auth/controller"
	"github.com/FrosTiK-SD/auth/util"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllDomains(ctx *gin.Context) {
	noCache := util.GetNoCache(ctx)

	domains, err := controller.GetAllDomains(h.MongikClient, noCache)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   constants.ERROR_MONGO_ERROR,
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": domains,
	})
}
