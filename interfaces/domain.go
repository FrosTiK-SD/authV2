package interfaces

import "github.com/FrosTiK-SD/auth/model"

type BatchCreateDomainRequest struct {
	Domains []model.Domain `json:"domains" bson:"domains" binding:"required"`
}

type UpdateDomainRequest struct {
	// The form key is "domains" (plural)
	Domain model.Domain `json:"domains" bson:"domain" binding:"required"`
}
