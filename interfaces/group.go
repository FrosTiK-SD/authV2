package interfaces

import (
	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/models/company"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssignRequest struct {
	Action constants.Action     `json:"action" bson:"action"`
	Groups []primitive.ObjectID `json:"groups" bson:"groups"`
	Roles  []string             `json:"roles" bson:"roles"`
}

type BatchCreateGroupRequest struct {
	Groups []company.Group `json:"group" bson:"group"`
}

type BatchDeleteGroupRequest struct {
	Groups []primitive.ObjectID `json:"groups" bson:"groups"`
}
