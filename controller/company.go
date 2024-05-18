package controller

import (
	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/auth/model"
	db "github.com/FrosTiK-SD/mongik/db"
	mongikModels "github.com/FrosTiK-SD/mongik/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllCompanies(mongikClient *mongikModels.Mongik, noCache bool) ([]model.Company, error) {
	companies, err := db.Aggregate[model.Company](mongikClient, constants.DB, constants.COLLECTION_COMPANY, []bson.M{}, noCache)

	return companies, err
}
