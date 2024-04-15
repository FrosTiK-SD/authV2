package interfaces

import (
	"github.com/FrosTiK-SD/auth/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssignRequest struct {
	Action constants.Action     `json:"action" bson:"action"`
	Groups []primitive.ObjectID `json:"groups" bson:"groups"`
	Roles  []string             `json:"roles" bson:"roles"`
}
