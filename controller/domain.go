package controller

import (
	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/auth/model"
	db "github.com/FrosTiK-SD/mongik/db"
	mongikModels "github.com/FrosTiK-SD/mongik/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllDomains(mongoClient *mongikModels.Mongik, noCache bool) ([]model.DomainPopulated, error) {
	domains, err := db.Aggregate[model.DomainPopulated](mongoClient, constants.DB, constants.COLLECTION_DOMAIN, []bson.M{{
		"$lookup": bson.M{
			"from":         constants.COLLECTION_STUDENT,
			"localField":   "assignedTo",
			"foreignField": "_id",
			"as":           "assignedTo",
		},
	}}, noCache)

	return domains, err
}
