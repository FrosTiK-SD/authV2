package model

import (
	studentModel "github.com/FrosTiK-SD/models/student"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Group struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Roles []string           `json:"roles" bson:"roles"`
}

type StudentPopulated struct {
	studentModel.Student
	GroupDetails []Group `json:"groups" bson:"groups"`
}
