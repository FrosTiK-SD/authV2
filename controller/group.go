package controller

import (
	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/models/company"
	db "github.com/FrosTiK-SD/mongik/db"
	mongikModels "github.com/FrosTiK-SD/mongik/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllGroups(mongikClient *mongikModels.Mongik, noCache bool) (*[]company.Group, error) {
	groups, err := db.Aggregate[company.Group](mongikClient, constants.DB, constants.COLLECTION_GROUP, []bson.M{}, noCache)

	return &groups, err
}
