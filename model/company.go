package model

import (
	"github.com/FrosTiK-SD/models/company"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	company.Company

	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
}
