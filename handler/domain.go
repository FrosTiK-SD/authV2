package handler

import (
	"net/http"

	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/auth/controller"
	"github.com/FrosTiK-SD/auth/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (h *Handler) GetDomainById(ctx *gin.Context) {
	noCache := util.GetNoCache(ctx)
	_id, err := primitive.ObjectIDFromHex(ctx.GetHeader("id"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   constants.ERROR_INCORRENT_BODY,
			"message": "Invalid Id",
		})
		return
	}

	domain, err := controller.GetDomainById(h.MongikClient, _id, noCache)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   constants.ERROR_MONGO_ERROR,
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": domain,
	})
}
