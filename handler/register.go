package handler

import (
	"github.com/FrosTiK-SD/authV2/interfaces"
	constant "github.com/FrosTiK-SD/models/constant"
	student "github.com/FrosTiK-SD/models/student"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (h *Handler) RegisterNewStudent(ctx *gin.Context) (interfaces.StudentProtoInterface, error) {
	student := interfaces.StudentProtoInterface{}

	if errBinding := ctx.ShouldBindBodyWith(&student, binding.JSON); errBinding != nil {
		return interfaces.StudentProtoInterface{}, errBinding
	}
}
