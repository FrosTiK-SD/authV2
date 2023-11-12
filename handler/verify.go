package handler

import (
	"fmt"

	"frostik.com/auth/constants"
	"frostik.com/auth/controller"
	"github.com/allegro/bigcache/v3"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	MongoClient     *mongo.Client
	UserCacheClient *bigcache.BigCache
}

func (h *Handler) HandlerVerifyIdToken(ctx *gin.Context) {
	idToken := ctx.GetHeader("token")
	fmt.Println(ctx.GetHeader("cache-control"))
	noCache := false
	if ctx.GetHeader("cache-control") == constants.NO_CACHE {
		noCache = true
	}

	email, err := controller.VerifyToken(h.UserCacheClient, idToken, noCache)

	if err != nil {
		ctx.JSON(200, gin.H{
			"student": nil,
			"error":   err,
		})
	} else {
		student := controller.GetStudentByEmail(ctx, h.MongoClient, h.UserCacheClient, email, noCache)
		ctx.JSON(200, gin.H{
			"data":  student,
			"error": nil,
		})
	}
}

func (h *Handler) InvalidateCache(ctx *gin.Context) {
	h.UserCacheClient.Delete("GCP_JWKS")
	ctx.JSON(200, gin.H{
		"message": "Successfully invalidated cache",
	})
}
