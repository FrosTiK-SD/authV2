package controller

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/auth/model"
	"github.com/FrosTiK-SD/auth/util"
	"github.com/FrosTiK-SD/models/group"
	db "github.com/FrosTiK-SD/mongik/db"
	models "github.com/FrosTiK-SD/mongik/models"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

type VerificationRequest struct {
    Branch string `json:"branch"`
    Course string `json:"course"`
    Batch  string `json:"batch"`
}

func startVerification(mongikClient *models.Mongik, c *gin.Context) {
    var req VerificationRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := processVerification(mongikClient, req, c); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process verification"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Verification started successfully"})
}



func processVerification(mongikClient *models.Mongik, req VerificationRequest, ctx *gin.Context) error {

	groupRoleName := "verification_" + req.Batch + "_" + req.Course + "_" + req.Branch
	filter := bson.M{"roles": groupRoleName}
	foundGroups, err := db.Aggregate[group.Group](mongikClient,constants.DB, constants.COLLECTION_GROUP, []bson.M{{"$match": filter}}, false)
	if err != nil {
		fmt.Println("Error checking for existing groups: ", err)
		return nil
	}

	var groupId primitive.ObjectID
	if len(foundGroups) == 0 {
		newGroup := group.Group{
			ID:    primitive.NewObjectID(),
			Name:  "Resume Verification Goup",
			Roles: []string{groupRoleName},
		}
		
		insertResult, err := db.InsertOne[group.Group](mongikClient,constants.DB, constants.COLLECTION_GROUP, newGroup)
		if err != nil {
			fmt.Println("Error creating new group: ", err)
			return nil
		}
		groupId = insertResult.InsertedID.(primitive.ObjectID) 
	} else {
		groupId = foundGroups[0].ID
	}


	pipeline := []bson.M{
		{
			"$match": bson.M{
				"roles": "TPR", 
			},
		},
			{
				"$project": bson.M{
					"_id": 1,
					"name": 1,
				},
			},
	}
	
	groupsWithRoleTPR, err := db.Aggregate[group.Group](mongikClient, constants.DB, constants.COLLECTION_GROUP, pipeline, false)
	if err != nil {
		fmt.Println("Error fetching groups with TPR role: ", err)
		return nil
	} 
	pipeline2 := []bson.M{
		{
			"$match": bson.M{
				"groups": bson.M{"$in": groupsWithRoleTPR},
				"batch.endYear": req.Batch, 
            	"department":req.Branch, 
            	"course": req.Course, 
			},
		},
		{
			"$project": bson.M{
				"_id": 1,
				"email": 1,
				"firstName": 1,
				"lastName": 1,
				"groups": 1,
			},
		},
	}

	studentTprs, err := db.Aggregate[model.StudentPopulated](mongikClient,constants.DB, constants.COLLECTION_GROUP, pipeline2, false)
	if err != nil {
		fmt.Println("Error fetching students in TPR groups: ", err)
		return nil
	}

	for _, student := range studentTprs {
		updateFilter := bson.M{"_id": student.Id}
		update := bson.M{"$addToSet": bson.M{"groups": groupId}}
		_, err := mongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_STUDENT).UpdateOne(context.Background(), updateFilter, update)
		if err != nil {
			fmt.Println("Error adding group to student: ", err)
		}
	}

	return nil
}