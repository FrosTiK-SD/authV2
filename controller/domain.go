package controller

import (
	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/auth/model"
	db "github.com/FrosTiK-SD/mongik/db"
	mongikModels "github.com/FrosTiK-SD/mongik/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetDomainById(mongoClient *mongikModels.Mongik, _id primitive.ObjectID, noCache bool) (*model.DomainPopulated, error) {
	domain, err := db.AggregateOne[model.DomainPopulated](mongoClient, constants.DB, constants.COLLECTION_DOMAIN, []bson.M{{
		"$match": bson.M{
			"_id": _id,
		},
	}, {
		"$lookup": bson.M{
			"from":         constants.COLLECTION_STUDENT,
			"localField":   "assignedTo",
			"foreignField": "_id",
			"as":           "assignedTo",
		},
	}}, noCache)

	return &domain, err
}
