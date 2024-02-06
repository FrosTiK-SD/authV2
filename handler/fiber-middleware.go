package handler

import (
	"errors"
	"net/http"

	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/auth/interfaces"
	"github.com/FrosTiK-SD/auth/util"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

// For Gin based middlewares
func (h *Handler) FiberVerifyStudent(ctx *fiber.Ctx) error {
	token := ctx.Get("token", "")

	// Create a new session
	currentHandler := Handler{
		MongikClient: h.MongikClient,
		JwkSet:       h.JwkSet,
		Session:      &Session{},
		Config: Config{
			Mode: MIDDLEWARE,
		},
	}

	context := gin.Context{
		Request: &http.Request{
			Header: http.Header{
				"token": []string{token},
			},
		},
	}

	currentHandler.HandlerVerifyStudentIdToken(&context)
	student := currentHandler.Session.Student

	if student != nil {
		ctx.Locals(constants.SESSION, student)
		ctx.Next()
	} else {
		return currentHandler.Session.Error
	}

	return nil
}

func (h *RoleCheckerHandler) FiberVerifyRole(ctx *fiber.Ctx) error {
	entity := ctx.Locals(constants.SESSION)
	var entityGroups *interfaces.Groups
	entityBytes, err := json.Marshal(entity)
	if err != nil {
		return err
	}
	err = json.Unmarshal(entityBytes, &entityGroups)
	if err != nil {
		return err
	}
	if !util.CheckRoleExists(&entityGroups.Groups, h.Role) {
		return errors.New(constants.ERROR_ROLE_CHECK_FAILED)
	}

	ctx.Next()
	return nil
}
