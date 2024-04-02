package model

import (
	group "github.com/FrosTiK-SD/models/company"
	studentModel "github.com/FrosTiK-SD/models/student"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudentPopulated struct {
	studentModel.Student
	GroupDetails []group.Group `json:"groups" bson:"groups"`
}

type Batch struct {
	StartYear int `json:"startYear" bson:"startYear"`
	EndYear   int `json:"endYear" bson:"endYear"`
}

type StudentMini struct {
	Id             primitive.ObjectID `json:"_id" bson:"_id"`
	Batch          *Batch             `json:"batch" bson:"batch"`
	RollNo         int                `json:"rollNo" bson:"rollNo"`
	InstituteEmail string             `json:"email" bson:"email"`
	Department     string             `json:"department" bson:"department"`
	Course         *string            `json:"course" bson:"course"`
	Specialisation *string            `json:"specialisation" bson:"specialisation"`
	FirstName      string             `json:"firstName" bson:"firstName"`
	MiddleName     *string            `json:"middleName" bson:"middleName"`
	LastName       *string            `json:"lastName" bson:"lastName"`
}
