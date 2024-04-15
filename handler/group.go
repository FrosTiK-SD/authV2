package handler

import (
	"net/http"

	"github.com/FrosTiK-SD/auth/controller"
	"github.com/FrosTiK-SD/auth/util"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllGroups(ctx *gin.Context) {
	noCache := util.GetNoCache(ctx)
	groups, err := controller.GetAllGroups(h.MongikClient, noCache)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err,
			"data":  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  groups,
		"error": nil,
	})
}
