package model

import (
	company "github.com/FrosTiK-SD/models/company"
)

type RecruiterModelPopulated struct {
	company.Recruiter
	GroupDetails []company.Group `json:"groups" bson:"groups"`
}
