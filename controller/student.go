package controller

import (
	"fmt"
	"time"

	"frostik.com/auth/constants"
	"frostik.com/auth/model"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetStudentByEmail(ctx *gin.Context, mongoClient *mongo.Client, cacheClient *cache.Cache, email *string) *model.Student {
	var student model.Student
	// Check if copy is there in the cache
	if err := cacheClient.Get(ctx, *email, &student); err == nil {
		fmt.Println("Successfully retreived user details from cache")
		return &student
	}

	// Query to DB
	fmt.Println("Queriying the DB for User Details")
	mongoClient.Database(constants.DB).Collection(constants.COLLECTION_STUDENT).FindOne(ctx, bson.M{
		"email": email,
	}).Decode(&student)

	// Set to redis
	if err := cacheClient.Set(&cache.Item{
		Ctx:   ctx,
		Key:   *email,
		Value: student,
		TTL:   time.Hour,
	}); err == nil {
		fmt.Println("Successfully set UserDetails in cache")
	}
	return &student
}
