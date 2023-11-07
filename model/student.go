package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID               primitive.ObjectID   `json:"_id" bson:"_id"`
	Batch            int                  `json:"batch"`
	RollNo           int                  `json:"rollNo"`
	FirstName        string               `json:"firstName"`
	LastName         string               `json:"lastName"`
	Department       string               `json:"department"`
	Course           string               `json:"course"`
	Email            string               `json:"email"`
	PersonalEmail    string               `json:"personalEmail"`
	LinkedIn         string               `json:"linkedIn"`
	Github           string               `json:"github"`
	MicrosoftTeams   string               `json:"microsoftTeams"`
	Mobile           int64                `json:"mobile"`
	Gender           string               `json:"gender"`
	Dob              string               `json:"dob"`
	PermanentAddress string               `json:"permanentAddress"`
	PresentAddress   string               `json:"presentAddress"`
	Category         string               `json:"category"`
	FatherName       string               `json:"fatherName"`
	FatherOccupation string               `json:"fatherOccupation"`
	MotherName       string               `json:"motherName"`
	MotherOccupation string               `json:"motherOccupation"`
	MotherTongue     string               `json:"motherTongue"`
	EducationGap     string               `json:"educationGap"`
	JeeRank          string               `json:"jeeRank"`
	Cgpa             float64              `json:"cgpa"`
	ActiveBacklogs   int                  `json:"activeBacklogs"`
	TotalBacklogs    int                  `json:"totalBacklogs"`
	XBoard           string               `json:"xBoard"`
	XYear            string               `json:"xYear"`
	XPercentage      int                  `json:"xPercentage"`
	XInstitute       string               `json:"xInstitute"`
	XiiBoard         string               `json:"xiiBoard"`
	XiiYear          string               `json:"xiiYear"`
	XiiPercentage    float64              `json:"xiiPercentage"`
	XiiInstitute     string               `json:"xiiInstitute"`
	SemesterOne      float64              `json:"semesterOne"`
	SemesterTwo      float64              `json:"semesterTwo"`
	SemesterThree    float64              `json:"semesterThree"`
	SemesterFour     float64              `json:"semesterFour"`
	SemesterFive     float64              `json:"semesterFive"`
	SemesterSix      int                  `json:"semesterSix"`
	Groups           []primitive.ObjectID `json:"groups"`
	UpdatedAt        string               `json:"updatedAt"`
}
