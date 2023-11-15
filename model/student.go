package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID               primitive.ObjectID   `json:"_id" bson:"_id" mapstructure:"_id"`
	Groups           []primitive.ObjectID `json:"groups" bson:"groups"`
	CompaniesAlloted []string             `json:"companiesAlloted" bson:"companiesAlloted"`
	Batch            int                  `json:"batch" bson:"batch"`
	RollNo           int                  `json:"rollNo" bson:"rollNo"`
	FirstName        string               `json:"firstName" bson:"firstName"`
	LastName         string               `json:"lastName" bson:"lastName"`
	Department       string               `json:"department" bson:"department"`
	Course           string               `json:"course" bson:"course"`
	Email            string               `json:"email" bson:"email"`
	PersonalEmail    string               `json:"personalEmail" bson:"personalEmail"`
	LinkedIn         string               `json:"linkedIn" bson:"linkedIn"`
	Github           string               `json:"github" bson:"github"`
	MicrosoftTeams   string               `json:"microsoftTeams" bson:"microsoftTeams"`
	Mobile           int64                `json:"mobile" bson:"mobile"`
	Gender           string               `json:"gender" bson:"gender"`
	Dob              string               `json:"dob" bson:"dob"`
	PermanentAddress string               `json:"permanentAddress" bson:"permanentAddress"`
	PresentAddress   string               `json:"presentAddress" bson:"presentAddress"`
	Category         string               `json:"category" bson:"category"`
	FatherName       string               `json:"fatherName" bson:"fatherName"`
	FatherOccupation string               `json:"fatherOccupation" bson:"fatherOccupation"`
	MotherName       string               `json:"motherName" bson:"motherName"`
	MotherOccupation string               `json:"motherOccupation" bson:"motherOccupation"`
	MotherTongue     string               `json:"motherTongue" bson:"motherTongue"`
	EducationGap     string               `json:"educationGap" bson:"educationGap"`
	JeeRank          string               `json:"jeeRank" bson:"jeeRank"`
	Cgpa             float64              `json:"cgpa" bson:"cgpa"`
	ActiveBacklogs   int                  `json:"activeBacklogs" bson:"activeBacklogs"`
	TotalBacklogs    int                  `json:"totalBacklogs" bson:"totalBacklogs"`
	XBoard           string               `json:"xBoard" bson:"xBoard"`
	XYear            string               `json:"xYear" bson:"xYear"`
	XPercentage      float64              `json:"xPercentage" bson:"xPercentage"`
	XInstitute       string               `json:"xInstitute" bson:"xInstitute"`
	XiiBoard         string               `json:"xiiBoard" bson:"xiiBoard"`
	XiiYear          string               `json:"xiiYear" bson:"xiiYear"`
	XiiPercentage    float64              `json:"xiiPercentage" bson:"xiiPercentage"`
	XiiInstitute     string               `json:"xiiInstitute" bson:"xiiInstitute"`
	SemesterOne      float64              `json:"semesterOne" bson:"semesterOne"`
	SemesterTwo      float64              `json:"semesterTwo" bson:"semesterTwo"`
	SemesterThree    float64              `json:"semesterThree" bson:"semesterThree"`
	SemesterFour     float64              `json:"semesterFour" bson:"semesterFour"`
	SemesterFive     float64              `json:"semesterFive" bson:"semesterFive"`
	SemesterSix      float64              `json:"semesterSix" bson:"semesterSix"`
	UpdatedAt        string               `json:"updatedAt" bson:"updatedAt"`
}

type Group struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Roles []string           `json:"roles" bson:"roles"`
}

type StudentPopulated struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	Batch            int                `json:"batch" bson:"batch"`
	RollNo           int                `json:"rollNo" bson:"rollNo"`
	FirstName        string             `json:"firstName" bson:"firstName"`
	LastName         string             `json:"lastName" bson:"lastName"`
	Department       string             `json:"department" bson:"department"`
	Course           string             `json:"course" bson:"course"`
	Email            string             `json:"email" bson:"email"`
	PersonalEmail    string             `json:"personalEmail" bson:"personalEmail"`
	LinkedIn         string             `json:"linkedIn" bson:"linkedIn"`
	Github           string             `json:"github" bson:"github"`
	MicrosoftTeams   string             `json:"microsoftTeams" bson:"microsoftTeams"`
	Mobile           int64              `json:"mobile" bson:"mobile"`
	Gender           string             `json:"gender" bson:"gender"`
	Dob              string             `json:"dob" bson:"dob"`
	PermanentAddress string             `json:"permanentAddress" bson:"permanentAddress"`
	PresentAddress   string             `json:"presentAddress" bson:"presentAddress"`
	Category         string             `json:"category" bson:"category"`
	FatherName       string             `json:"fatherName" bson:"fatherName"`
	FatherOccupation string             `json:"fatherOccupation" bson:"fatherOccupation"`
	MotherName       string             `json:"motherName" bson:"motherName"`
	MotherOccupation string             `json:"motherOccupation" bson:"motherOccupation"`
	MotherTongue     string             `json:"motherTongue" bson:"motherTongue"`
	EducationGap     string             `json:"educationGap" bson:"educationGap"`
	JeeRank          string             `json:"jeeRank" bson:"jeeRank"`
	Cgpa             float64            `json:"cgpa" bson:"cgpa"`
	ActiveBacklogs   int                `json:"activeBacklogs" bson:"activeBacklogs"`
	TotalBacklogs    int                `json:"totalBacklogs" bson:"totalBacklogs"`
	XBoard           string             `json:"xBoard" bson:"xBoard"`
	XYear            string             `json:"xYear" bson:"xYear"`
	XPercentage      float64            `json:"xPercentage" bson:"xPercentage"`
	XInstitute       string             `json:"xInstitute" bson:"xInstitute"`
	XiiBoard         string             `json:"xiiBoard" bson:"xiiBoard"`
	XiiYear          string             `json:"xiiYear" bson:"xiiYear"`
	XiiPercentage    float64            `json:"xiiPercentage" bson:"xiiPercentage"`
	XiiInstitute     string             `json:"xiiInstitute" bson:"xiiInstitute"`
	SemesterOne      float64            `json:"semesterOne" bson:"semesterOne"`
	SemesterTwo      float64            `json:"semesterTwo" bson:"semesterTwo"`
	SemesterThree    float64            `json:"semesterThree" bson:"semesterThree"`
	SemesterFour     float64            `json:"semesterFour" bson:"semesterFour"`
	SemesterFive     float64            `json:"semesterFive" bson:"semesterFive"`
	SemesterSix      float64            `json:"semesterSix" bson:"semesterSix"`
	Groups           []Group            `json:"groups" bson:"groups"`
	UpdatedAt        string             `json:"updatedAt" bson:"updatedAt"`
	CompaniesAlloted []string           `json:"companiesAlloted" bson:"companiesAlloted"`
}
