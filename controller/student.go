package controller

import (
	"encoding/json"
	"fmt"

	"frostik.com/auth/constants"
	"frostik.com/auth/model"
	"github.com/allegro/bigcache/v3"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetStudentByEmail(ctx *gin.Context, mongoClient *mongo.Client, cacheClient *bigcache.BigCache, email *string, noCache bool) *model.Student {
	var student model.Student
	var studentBytes []byte

	// Check if copy is there in the cache
	if !noCache {
		studentBytes, _ := cacheClient.Get(*email)
		fmt.Println(string(studentBytes))
		if err := json.Unmarshal(studentBytes, &student); err == nil {
			fmt.Println("Retreiving the student data from the cache")
			return &student
		}
	}

	// Query to DB
	fmt.Println("Queriying the DB for User Details")
	mongoClient.Database(constants.DB).Collection(constants.COLLECTION_STUDENT).FindOne(ctx, bson.M{
		"email": email,
	}).Decode(&student)

	// Set to bigCache
	studentBytes, _ = json.Marshal(student)
	if err := cacheClient.Set(*email, studentBytes); err == nil {
		fmt.Println("Successfully set UserDetails in cache")
	}
	return &student
}
