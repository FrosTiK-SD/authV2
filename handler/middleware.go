package handler

import (
	"github.com/FrosTiK-SD/auth/constants"
	"github.com/gin-gonic/gin"
)

// For Gin based middlewares
func (h *Handler) GinVerifyStudent(ctx *gin.Context) {

	// Create a new session
	currentHandler := Handler{
		MongikClient: h.MongikClient,
		JwkSet:       h.JwkSet,
		Session:      &Session{},
		Config: Config{
			Mode: MIDDLEWARE,
		},
	}

	currentHandler.HandlerVerifyStudentIdToken(ctx)
	student := currentHandler.Session.Student

	if student != nil {
		ctx.Set(constants.SESSION_STUDENT, student)
		ctx.Next()
	}
}
