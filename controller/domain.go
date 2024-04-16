package controller

import (
	"time"

	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/auth/model"
	"github.com/FrosTiK-SD/models/student"
	db "github.com/FrosTiK-SD/mongik/db"
	mongikModels "github.com/FrosTiK-SD/mongik/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllDomains(mongikClient *mongikModels.Mongik, noCache bool) ([]model.DomainPopulated, error) {
	domains, err := db.Aggregate[model.DomainPopulated](mongikClient, constants.DB, constants.COLLECTION_DOMAIN, []bson.M{{
		"$lookup": bson.M{
			"from":         constants.COLLECTION_STUDENT,
			"localField":   "assignedTo",
			"foreignField": "_id",
			"as":           "assignedTo",
		},
	}}, noCache)

	return domains, err
}

func GetDomainById(mongikClient *mongikModels.Mongik, _id primitive.ObjectID, noCache bool) (*model.DomainPopulated, error) {
	domain, err := db.AggregateOne[model.DomainPopulated](mongikClient, constants.DB, constants.COLLECTION_DOMAIN, []bson.M{{
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

func BatchCreateDomain(mongikClient *mongikModels.Mongik, domains []model.Domain) (*mongo.InsertManyResult, []*mongo.UpdateResult, []error) {

	for idx := range domains {
		domains[idx].CreatedAt = primitive.NewDateTimeFromTime(time.Now())
		domains[idx].UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	}

	domainResult, err := db.InsertMany[model.Domain](mongikClient, constants.DB, constants.COLLECTION_DOMAIN, domains)

	if err != nil {
		return domainResult, nil, []error{err}
	}

	var errors []error
	var studentResults []*mongo.UpdateResult

	for idx := range domains {

		studentResult, err := db.UpdateMany[student.Student](mongikClient, constants.DB, constants.COLLECTION_STUDENT, bson.M{
			"_id": bson.M{
				"$in": domains[idx].AssignedTo,
			},
		}, bson.M{
			"$addToSet": bson.M{
				"companiesAlloted": domains[idx].Domain,
			},
		})
		studentResults = append(studentResults, studentResult)
		if err != nil {
			errors = append(errors, err)
		}
	}

	return domainResult, studentResults, errors
}

func UpdateDomainById(mongikClient *mongikModels.Mongik, domainId primitive.ObjectID, updatedDomain *model.Domain) (*model.Domain, *mongo.UpdateResult, *mongo.UpdateResult, error) {

	oldDomain := db.FindOneAndUpdate[model.Domain](mongikClient, constants.DB, constants.COLLECTION_DOMAIN, bson.M{
		"_id": domainId,
	},
		bson.M{
			"$set": bson.M{
				"domain":      updatedDomain.Domain,
				"companyName": updatedDomain.CompanyName,
				"assignedTo":  updatedDomain.AssignedTo,

				"updatedAt": primitive.NewDateTimeFromTime(time.Now()),
			},
		},
	)

	removeDomainResult, err := db.UpdateMany[student.Student](mongikClient, constants.DB, constants.COLLECTION_STUDENT, bson.M{
		"_id": bson.M{
			"$in": oldDomain.AssignedTo,
		},
	}, bson.M{
		"$pull": bson.M{
			"companiesAlloted": oldDomain.Domain,
		},
	})

	if err != nil {
		return &oldDomain, removeDomainResult, nil, err
	}

	addDomainResult, err := db.UpdateMany[student.Student](mongikClient, constants.DB, constants.COLLECTION_STUDENT, bson.M{
		"_id": bson.M{
			"$in": updatedDomain.AssignedTo,
		},
	}, bson.M{
		"$addToSet": bson.M{
			"companiesAlloted": updatedDomain.Domain,
		},
	})

	return &oldDomain, removeDomainResult, addDomainResult, err
}
