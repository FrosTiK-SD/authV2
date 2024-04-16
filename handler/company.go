package handler

import (
	"net/http"

	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/auth/controller"
	"github.com/FrosTiK-SD/auth/util"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllCompanies(ctx *gin.Context) {
	noCache := util.GetNoCache(ctx)

	companies, err := controller.GetAllCompanies(h.MongikClient, noCache)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   constants.ERROR_MONGO_ERROR,
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": companies,
	})
}
