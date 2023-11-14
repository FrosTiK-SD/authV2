package handler

import (
	"frostik.com/auth/constants"
	"frostik.com/auth/controller"
	"github.com/allegro/bigcache/v3"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	MongoClient     *mongo.Client
	UserCacheClient *bigcache.BigCache
	JwkSet          *jwk.Set
}

func (h *Handler) HandlerVerifyStudentIdToken(ctx *gin.Context) {
	idToken := ctx.GetHeader("token")
	noCache := false
	if ctx.GetHeader("cache-control") == constants.NO_CACHE {
		noCache = true
	}

	email, err := controller.VerifyToken(h.UserCacheClient, idToken, h.JwkSet, noCache)

	if err != nil {
		ctx.JSON(200, gin.H{
			"student": nil,
			"error":   err,
		})
	} else {
		student, err := controller.GetUserByEmail(ctx, h.MongoClient, h.UserCacheClient, email, &constants.ROLE_STUDENT, noCache)
		ctx.JSON(200, gin.H{
			"data":  student,
			"error": err,
		})
	}
}

func (h *Handler) InvalidateCache(ctx *gin.Context) {
	h.UserCacheClient.Delete("GCP_JWKS")
	ctx.JSON(200, gin.H{
		"message": "Successfully invalidated cache",
	})
}
