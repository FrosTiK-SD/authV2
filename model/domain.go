package model

import (
	"github.com/FrosTiK-SD/models/student"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Domain struct {
	ID          primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Domain      string               `json:"domain" bson:"domain" binding:"required"`
	CompanyName string               `json:"companyName" bson:"companyName" binding:"required"`
	AssignedTo  []primitive.ObjectID `json:"assignedTo" bson:"assignedTo" binding:"required"`

	CreatedAt primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type DomainPopulated struct {
	Domain
	AssignedTo []student.Student `json:"assignedTo" bson:"assignedTo"`
}
