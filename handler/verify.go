package handler

import (
	"frostik.com/auth/controller"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	MongoClient     *mongo.Client
	UserRedis       *redis.Client
	UserCacheClient *cache.Cache
}

func (h *Handler) HandlerVerifyIdToken(ctx *gin.Context) {
	idToken := ctx.GetHeader("token")

	email, err := controller.VerifyToken(ctx, h.UserCacheClient, idToken)

	if err != nil {
		ctx.JSON(200, gin.H{
			"student": nil,
			"error":   err,
		})
	} else {
		student := controller.GetStudentByEmail(ctx, h.MongoClient, h.UserCacheClient, email)
		ctx.JSON(200, gin.H{
			"data":  student,
			"error": nil,
		})
	}
}

func (h *Handler) InvalidateCache(ctx *gin.Context) {
	h.UserCacheClient.Delete(ctx, "GCP_JWKS")
	ctx.JSON(200, gin.H{
		"message": "Successfully invalidated cache",
	})
}
