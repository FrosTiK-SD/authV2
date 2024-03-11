package handler

import (
	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/auth/controller"
	studentModel "github.com/FrosTiK-SD/models/student"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *Handler) HandlerUpdateStudentDetails(ctx *gin.Context) {
	studentCollection := h.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_STUDENT)
	updatedStudent := studentModel.Student{}
	if errBinding := ctx.ShouldBindBodyWith(&updatedStudent, binding.JSON); errBinding != nil {
		ctx.AbortWithStatusJSON(401, gin.H{"error": errBinding.Error()})
		return
	}

	filter := bson.M{"_id": h.Session.Student.Id, "email": h.Session.Student.InstituteEmail}
	currentStudent := studentModel.Student{}

	if errFind := studentCollection.FindOne(ctx, filter).Decode(&currentStudent); errFind != nil {
		ctx.AbortWithStatusJSON(401, gin.H{"error": errFind.Error()})
		return
	}

	controller.AssignUnVerifiedFields(&updatedStudent, &currentStudent)
	controller.InvalidateVerifiedFieldsOnChange(&updatedStudent, &currentStudent)

	if updateResult, errUpdate := studentCollection.ReplaceOne(ctx, filter, currentStudent); errUpdate != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": errUpdate.Error()})
		return
	} else {
		ctx.JSON(200, gin.H{"student": updateResult})

	}
}
