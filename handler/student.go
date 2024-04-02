package handler

import (
	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/auth/controller"
	"github.com/FrosTiK-SD/auth/interfaces"
	"github.com/gin-gonic/gin"
)

func (h *Handler) HandlerGetStudentsByRollNos(ctx *gin.Context) {
	h.GinVerifyStudent(ctx)

	noCache := false
	if ctx.GetHeader("cache-control") == constants.NO_CACHE {
		noCache = true
	}

	var requestBody interfaces.StudentsRollNoReq
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": constants.ERROR_INCORRENT_BODY,
		})
		return
	}

	students, err := controller.GetStudentsByRollNos(h.MongikClient, &requestBody, noCache)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"data":  nil,
			"error": err,
		})
	} else {
		ctx.JSON(200, gin.H{
			"data":  students,
			"error": nil,
		})
	}

}
