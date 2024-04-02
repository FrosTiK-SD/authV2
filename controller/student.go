package controller

import (
	"sort"
	"strings"

	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/auth/interfaces"
	"github.com/FrosTiK-SD/auth/model"
	"github.com/FrosTiK-SD/auth/util"
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

	// Gets the alias emails
	emailList := getAliasEmailList(*email)

	// Query to DB
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
	if !util.CheckRoleExists(&studentPopulated.GroupDetails, *role) {
		return nil, &constants.ERROR_NOT_A_STUDENT
	}

	return &studentPopulated, nil
}

func GetStudentsByRollNos(mongikClient *models.Mongik, request *interfaces.StudentsRollNoReq, noCache bool) (*[]model.StudentMini, error) {
	var studentsMini []model.StudentMini

	studentsMini, err := db.Find[model.StudentMini](mongikClient, constants.DB, constants.COLLECTION_STUDENT, bson.M{
		"rollNo": bson.M{"$in": request.RollNos},
	}, noCache)
	if err != nil {
		return nil, err
	}

	return &studentsMini, nil
}
