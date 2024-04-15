package controller

import (
	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/auth/interfaces"
	"github.com/FrosTiK-SD/models/company"
	db "github.com/FrosTiK-SD/mongik/db"
	mongikModels "github.com/FrosTiK-SD/mongik/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllGroups(mongikClient *mongikModels.Mongik, noCache bool) (*[]company.Group, error) {
	groups, err := db.Aggregate[company.Group](mongikClient, constants.DB, constants.COLLECTION_GROUP, []bson.M{}, noCache)

	return &groups, err
}

func BatchCreateGroup(mongikClient *mongikModels.Mongik, groups []company.Group) (*mongo.InsertManyResult, error) {

	for idx, _ := range groups {
		groups[idx].ID = primitive.NewObjectID()
	}

	insertResult, err := db.InsertMany(mongikClient, constants.DB, constants.COLLECTION_GROUP, groups)

	return insertResult, err
}

func BatchEditGroup(mongikClient *mongikModels.Mongik, assignRequests []interfaces.AssignRequest, noCache bool) (*[]*mongo.UpdateResult, *[]*mongo.UpdateResult, *[]error) {
	var addList, removeList []*mongo.UpdateResult
	var errors []error
	for _, request := range assignRequests {
		switch request.Action {
		case constants.ACTION_PUSH:
			addResult, err := db.UpdateMany[company.Group](mongikClient, constants.DB, constants.COLLECTION_GROUP, bson.M{
				"_id": bson.M{
					"$in": request.Groups,
				},
			}, bson.M{
				"$addToSet": bson.M{
					"roles": bson.M{
						"$each": request.Roles,
					},
				},
			})
			addList = append(addList, addResult)
			if err != nil {
				errors = append(errors, err)
			}
		case constants.ACTION_PULL:
			removeResult, err := db.UpdateMany[company.Group](mongikClient, constants.DB, constants.COLLECTION_GROUP, bson.M{
				"_id": bson.M{
					"$in": request.Groups,
				},
			}, bson.M{
				"$pull": bson.M{
					"roles": bson.M{
						"$in": request.Roles,
					},
				},
			})
			removeList = append(removeList, removeResult)
			if err != nil {
				errors = append(errors, err)
			}
		}

	}
	return &addList, &removeList, &errors
}
