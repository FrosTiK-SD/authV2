package handler

import (
	"context"

	"firebase.google.com/go/auth"
	"frostik.com/auth/controller"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	Auth        *auth.Client
	MongoClient *mongo.Client
}

func (h *Handler) HandlerVerifyIdToken(ctx *gin.Context) {
	idToken := ctx.GetHeader("token")

	token, err := h.Auth.VerifyIDToken(ctx, idToken)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "UNAUTHENTICATED",
		})
		return
	}
	firebaseUser, err := h.Auth.GetUser(context.Background(), token.UID)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "User does not exist",
		})
		return
	}

	email := firebaseUser.Email
	user := controller.GetStudentByEmail(ctx, h.MongoClient, &email)

	ctx.JSON(200, gin.H{
		"message": "Token verified successfully",
		"token":   token,
		"email":   firebaseUser,
		"user":    user,
		"data": gin.H{
			"data": user,
		},
	})
}
