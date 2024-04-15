package handler

import (
	"os"

	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/auth/controller"
	"github.com/FrosTiK-SD/auth/interfaces"
	"github.com/FrosTiK-SD/auth/model"
	"github.com/FrosTiK-SD/auth/util"
	"github.com/FrosTiK-SD/models/constant"
	studentModel "github.com/FrosTiK-SD/models/student"
	db "github.com/FrosTiK-SD/mongik/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetStudentRoleObjectID() primitive.ObjectID {
	if objID, err := primitive.ObjectIDFromHex(os.Getenv(constants.ENV_STUDENT_GROUP_OBJ_ID)); err != nil {
		return primitive.NilObjectID
	} else {
		return objID
	}
}

func (h *Handler) HandlerUpdateStudentDetails(ctx *gin.Context) {
	studentCollection := h.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_STUDENT)

	student, exists := ctx.Get(constants.SESSION)
	if !exists {
		ctx.AbortWithStatusJSON(401, gin.H{"error": "Cant get student"})
		return
	}

	studentPopulated := student.(*model.StudentPopulated)

	var updatedStudent studentModel.Student
	if errBinding := ctx.ShouldBindJSON(&updatedStudent); errBinding != nil {
		ctx.AbortWithStatusJSON(401, gin.H{"error": errBinding.Error()})
		return
	}

	filter := bson.M{"_id": studentPopulated.Id, "email": studentPopulated.InstituteEmail}

	var currentStudent studentModel.Student
	if errFind := studentCollection.FindOne(ctx, filter).Decode(&currentStudent); errFind != nil {
		ctx.AbortWithStatusJSON(401, gin.H{"error": errFind.Error()})
		return
	}

	controller.AssignUnVerifiedFields(&updatedStudent, &currentStudent)
	controller.InvalidateVerifiedFieldsOnChange(&updatedStudent, &currentStudent)

	if updateResult, errUpdate := db.ReplaceOne(h.MongikClient, constants.DB, constants.COLLECTION_STUDENT, filter, &currentStudent); errUpdate != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": errUpdate.Error()})
		return
	} else {
		ctx.JSON(200, gin.H{"student": updateResult})
	}
}

func (h *Handler) HandlerRegisterStudentDetails(ctx *gin.Context) {
	idToken := ctx.GetHeader("token")
	newStudentDetails := interfaces.StudentRegistration{}

	if errBinding := ctx.BindJSON(&newStudentDetails); errBinding != nil {
		return
	}

	if email, _, errVerify := controller.VerifyToken(h.MongikClient.CacheClient, idToken, h.JwkSet, true); errVerify != nil {
		ctx.AbortWithStatusJSON(401, gin.H{"error": errVerify})
		return
	} else {
		if !util.CheckValidInstituteEmail(*email) {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "not a valid institute email"})
			return
		}
		newStudentDetails.InstituteEmail = *email
	}

	newStudent := studentModel.Student{
		Groups:         []primitive.ObjectID{GetStudentRoleObjectID()},
		Id:             primitive.NewObjectID(),
		Batch:          &newStudentDetails.Batch,
		RollNo:         newStudentDetails.RollNo,
		InstituteEmail: newStudentDetails.InstituteEmail,
		Department:     newStudentDetails.Department,
		Course:         (*constant.Course)(&newStudentDetails.Course),
		Specialisation: newStudentDetails.Specialisation,
		FirstName:      newStudentDetails.FirstName,
		MiddleName:     newStudentDetails.MiddleName,
		LastName:       newStudentDetails.LastName,
		PersonalEmail:  newStudentDetails.PersonalEmail,
		Mobile:         newStudentDetails.Mobile,
		Gender:         newStudentDetails.Gender,
	}

	if result, err := db.InsertOne(h.MongikClient, constants.DB, constants.COLLECTION_STUDENT, newStudent); err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(200, gin.H{"student": newStudent, "logs": result})
		return
	}
}

func (h *Handler) HandlerGetStudentProfile(ctx *gin.Context) {
	student, exists := ctx.Get(constants.SESSION)
	if !exists {
		ctx.AbortWithStatusJSON(401, gin.H{"error": "Cant get student"})
		return
	}

	studentPopulated := student.(*model.StudentPopulated)
	studentProfile := controller.MapStudentToStudentProfile(studentPopulated)

	ctx.JSON(200, gin.H{"profile": studentProfile})
}
