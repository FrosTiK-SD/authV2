package interfaces

import (
	group "github.com/FrosTiK-SD/models/company"
)

type Groups struct {
	Groups []group.Group `json:"groups" bson:"groups"`
}
