package interfaces

import "github.com/FrosTiK-SD/auth/model"

type BatchCreateDomainRequest struct {
	Domains []model.Domain `json:"domains" bson:"domains" binding:"required"`
}
