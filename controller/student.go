package controller

import (
	"encoding/json"
	"fmt"
	"strings"

	"frostik.com/auth/constants"
	"frostik.com/auth/db"
	"frostik.com/auth/mapper"
	"frostik.com/auth/model"
	"frostik.com/auth/util"
	"github.com/allegro/bigcache/v3"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getAliasEmailList(email string) []string {
	var aliasEmailList []string
	aliasEmailList = append(aliasEmailList, email)
	aliasEmailList = append(aliasEmailList, strings.ReplaceAll(email, "iitbhu.ac.in", "itbhu.ac.in"))
	aliasEmailList = append(aliasEmailList, strings.ReplaceAll(email, "itbhu.ac.in", "iitbhu.ac.in"))
	return aliasEmailList
}

func GetUserByEmail(ctx *gin.Context, mongoClient *mongo.Client, cacheClient *bigcache.BigCache, email *string, role *string, noCache bool) (*model.StudentPopulated, *string) {
	var student model.Student
	var studentPopulated model.StudentPopulated

	// Check if copy is there in the cache
	if !noCache {
		studentBytes, _ := cacheClient.Get(*email)
		if err := json.Unmarshal(studentBytes, &studentPopulated); err == nil {
			fmt.Println("Retreiving the student data from the cache")
			return &studentPopulated, nil
		}
	}

	// Gets the alias emails
	emailList := getAliasEmailList(*email)

	// Query to DB
	fmt.Println("Queriying the DB for User Details")
	db.FindOne[model.Student](ctx, mongoClient, cacheClient, constants.COLLECTION_STUDENT, bson.M{
		"email": bson.M{"$in": emailList},
	}, &student, noCache)
	studentPopulated = mapper.TransformStudentToStudentPopulated(student)

	var groupIds = []primitive.ObjectID{}
	var groupDetails = []model.Group{}
	for _, id := range student.Groups {
		groupIds = append(groupIds, id)
	}

	groupDetails, _ = db.Find[model.Group](ctx, mongoClient, cacheClient, constants.COLLECTION_GROUP, bson.M{
		"_id": bson.M{"$in": groupIds},
	}, noCache)
	studentPopulated.Groups = groupDetails

	// Now check if it is actually a student by the ROLES
	if !util.CheckRoleExists(&groupDetails, *role) {
		return nil, &constants.ERROR_NOT_A_STUDENT
	}

	// Set to bigCache
	studentBytes, _ := json.Marshal(studentPopulated)
	if err := cacheClient.Set(*email, studentBytes); err == nil {
		fmt.Println("Successfully set UserDetails in cache")
	}
	return &studentPopulated, nil
}
