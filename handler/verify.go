package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/auth/controller"

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
		if h.Config.Mode == MIDDLEWARE {
			h.Session.Error = errors.New(*err)
		}
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
			if h.Config.Mode == MIDDLEWARE {
				h.Session.Error = errors.New(*err)
			}
			ctx.JSON(200, gin.H{
				"data":   student,
				"error":  err,
				"expire": exp,
			})
		}
	}
}

func (h *Handler) HandlerVerifyRecruiterIdToken(ctx *gin.Context) {

	handleError := func(err *string, exp *time.Time) {
		// TODO consider implementing a middle for it
		// if h.Config.Mode == MIDDLEWARE {
		// 	h.Session.Error = errors.New(*err)
		// }
		ctx.JSON(http.StatusOK, gin.H{
			"data":   nil,
			"expire": exp,
			"error":  err,
		})
	}

	idToken := ctx.GetHeader("token")
	noCache := false
	if ctx.GetHeader("cache-control") == constants.NO_CACHE {
		noCache = true
	}
	email, exp, err := controller.VerifyToken(h.MongikClient.CacheClient, idToken, h.JwkSet, noCache)
	fmt.Println("Email ", *email)
	if err != nil {
		handleError(err, exp)
		return
	}

	recruiter, err := controller.GetRecruiterByEmail(h.MongikClient, email, &constants.ROLE_RECRUITER, noCache)

	if err != nil {
		handleError(err, exp)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":   recruiter,
		"expire": exp,
		"error":  nil,
	})
}

func (h *Handler) InvalidateCache(ctx *gin.Context) {
	h.MongikClient.CacheClient.Delete("GCP_JWKS")
	ctx.JSON(200, gin.H{
		"message": "Successfully invalidated cache",
	})
}
