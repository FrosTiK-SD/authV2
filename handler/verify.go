package handler

import (
	"encoding/json"

	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/auth/controller"
	"github.com/FrosTiK-SD/auth/interfaces"
	"github.com/FrosTiK-SD/auth/util"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HandlerVerifyStudentIdToken(ctx *gin.Context) {
	idToken := ctx.GetHeader("token")
	noCache := false
	if ctx.GetHeader("cache-control") == constants.NO_CACHE {
		noCache = true
	}

	email, exp, err := controller.VerifyToken(h.MongikClient.CacheClient, idToken, h.JwkSet, noCache)

	if err != nil {
		ctx.JSON(200, gin.H{
			"student": nil,
			"expire":  exp,
			"error":   err,
		})
	} else {
		student, err := controller.GetUserByEmail(h.MongikClient, email, &constants.ROLE_STUDENT, noCache)
		if h.Config.Mode == MIDDLEWARE {
			h.Session.Student = student
		} else {
			ctx.JSON(200, gin.H{
				"data":   student,
				"error":  err,
				"expire": exp,
			})
		}
	}

}

func (h *Handler) InvalidateCache(ctx *gin.Context) {
	h.MongikClient.CacheClient.Delete("GCP_JWKS")
	ctx.JSON(200, gin.H{
		"message": "Successfully invalidated cache",
	})
}

func (h *RoleCheckerHandler) CheckRoleInGroup(ctx *gin.Context) {
	entity, exists := ctx.Get(constants.SESSION)
	if exists != true {
		ctx.JSON(200, gin.H{
			"role_exists": false,
			"error":       "Entity does not exist",
		})
		ctx.Abort()
	} else {
		var entityGroups *interfaces.Groups
		entityBytes, _ := json.Marshal(entity)
		json.Unmarshal(entityBytes, &entityGroups)
		if !util.CheckRoleExists(&entityGroups.Groups, h.Role) {
			ctx.JSON(200, gin.H{
				"role_exists": false,
			})
			ctx.Abort()
		} else {
			ctx.Next()
		}
	}
}