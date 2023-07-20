package controller

import (
	"frostik.com/auth/constants"
	"frostik.com/auth/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetStudentByEmail(ctx *gin.Context, mongoClient *mongo.Client, email *string) *model.Student {
	var student model.Student
	mongoClient.Database(constants.DB).Collection(constants.COLLECTION_STUDENT).FindOne(ctx, bson.M{
		"email": email,
	}).Decode(&student)
	return &student
}
