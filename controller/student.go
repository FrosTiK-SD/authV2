package controller

import (
	"fmt"
	"sort"
	"strings"

	"frostik.com/auth/constants"
	"frostik.com/auth/model"
	"frostik.com/auth/util"
	db "github.com/FrosTiK-SD/mongik/db"
	models "github.com/FrosTiK-SD/mongik/models"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func getAliasEmailList(email string) []string {
	var aliasEmailList []string
	aliasEmailList = append(aliasEmailList, email)
	aliasEmailList = append(aliasEmailList, strings.ReplaceAll(email, "iitbhu.ac.in", "itbhu.ac.in"))
	aliasEmailList = append(aliasEmailList, strings.ReplaceAll(email, "itbhu.ac.in", "iitbhu.ac.in"))
	sort.Strings(aliasEmailList)
	return aliasEmailList
}

func GetUserByEmail(mongikClient *models.Mongik, email *string, role *string, noCache bool) (*model.StudentPopulated, *string) {
	var studentPopulated model.StudentPopulated

	// Check if copy is there in the cache
	if !noCache {
		studentBytes, _ := mongikClient.CacheClient.Get(*email)
		if err := json.Unmarshal(studentBytes, &studentPopulated); err == nil {
			fmt.Println("Retreiving the student data from the cache")
			return &studentPopulated, nil
		}
	}

	// Gets the alias emails
	emailList := getAliasEmailList(*email)

	// Query to DB
	fmt.Println("Queriying the DB for User Details")
	studentPopulated, _ = db.AggregateOne[model.StudentPopulated](mongikClient, constants.DB, constants.COLLECTION_STUDENT, []bson.M{{
		"$match": bson.M{"email": bson.M{"$in": emailList}},
	}, {
		"$lookup": bson.M{
			"from":         constants.COLLECTION_GROUP,
			"localField":   "groups",
			"foreignField": "_id",
			"as":           "groups",
		},
	}}, noCache)

	// Now check if it is actually a student by the ROLES
	if !util.CheckRoleExists(&studentPopulated.Groups, *role) {
		return nil, &constants.ERROR_NOT_A_STUDENT
	}

	// Set to bigCache
	studentBytes, _ := json.Marshal(studentPopulated)
	if err := mongikClient.CacheClient.Set(*email, studentBytes); err == nil {
		fmt.Println("Successfully set UserDetails in cache")
	}
	return &studentPopulated, nil
}
