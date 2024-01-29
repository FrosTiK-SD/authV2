package interfaces

import "github.com/FrosTiK-SD/auth/model"

type Groups struct {
	Groups []model.Group `json:"groups" bson:"groups"`
}