package model

import (
	group "github.com/FrosTiK-SD/models/company"
	studentModel "github.com/FrosTiK-SD/models/student"
)

type StudentPopulated struct {
	studentModel.Student
	GroupDetails []group.Group `json:"groups" bson:"groups"`
}
