package interfaces

import (
	constant "github.com/FrosTiK-SD/models/constant"
	student "github.com/FrosTiK-SD/models/student"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Batch struct {
	StartYear int `json:"startYear" bson:"startYear"`
	EndYear   int `json:"endYear" bson:"endYear"`
}

type StudentProtoInterface struct {
	ID               primitive.ObjectID   `json:"_id" bson:"_id"`
	Groups           []primitive.ObjectID `json:"groups" bson:"groups"`
	CompaniesAlloted []string             `json:"companiesAlloted" bson:"companiesAlloted"`

	Batch          Batch           `json:"batch" bson:"batch"`
	RollNo         int             `json:"rollNo" bson:"rollNo"`
	InstituteEmail string          `json:"email" bson:"email"`
	Department     string          `json:"department" bson:"department"`
	Course         constant.Course `json:"course" bson:"course"`

	// metadata
	StructVersion int                `json:"version,omitempty" bson:"version,omitempty"`
	UpdatedAt     primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
	CreatedAt     primitive.DateTime `json:"createdAt" bson:"createdAt"`
}
