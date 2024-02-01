package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OldStudent struct {
	ID               primitive.ObjectID   `json:"_id" bson:"_id"`
	Groups           []primitive.ObjectID `json:"groups" bson:"groups"`
	CompaniesAlloted []string             `json:"companiesAlloted" bson:"companiesAlloted"`
	Batch            int                  `json:"batch" bson:"batch"`
	RollNo           int                  `json:"rollNo" bson:"rollNo"`
	FirstName        string               `json:"firstName" bson:"firstName"`
	LastName         *string              `json:"lastName" bson:"lastName"`
	Department       string               `json:"department" bson:"department"`
	Course           string               `json:"course" bson:"course"`
	Email            string               `json:"email" bson:"email"`
	PersonalEmail    string               `json:"personalEmail" bson:"personalEmail"`
	LinkedIn         string               `json:"linkedIn" bson:"linkedIn"`
	Github           string               `json:"github" bson:"github"`
	MicrosoftTeams   string               `json:"microsoftTeams" bson:"microsoftTeams"`
	Kaggle           string               `json:"kaggle" bson:"kaggle"`
	Kaggel           string               `json:"kaggel" bson:"kaggel"`
	Skype            string               `json:"skypeId" bson:"skypeId"`
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
	Cgpa             *float64             `json:"cgpa" bson:"cgpa"`
	ActiveBacklogs   *int                 `json:"activeBacklogs" bson:"activeBacklogs"`
	TotalBacklogs    *int                 `json:"totalBacklogs" bson:"totalBacklogs"`
	XBoard           string               `json:"xBoard" bson:"xBoard"`
	XPercentage      float64              `json:"xPercentage" bson:"xPercentage"`
	XInstitute       string               `json:"xInstitute" bson:"xInstitute"`
	XiiBoard         string               `json:"xiiBoard" bson:"xiiBoard"`
	XiiPercentage    float64              `json:"xiiPercentage" bson:"xiiPercentage"`
	XiiInstitute     string               `json:"xiiInstitute" bson:"xiiInstitute"`
	PgIn             string               `json:"pgIn" bson:"pgIn"`
	PgCollege        string               `json:"pgCollege" bson:"pgCollege"`
	PgScore          float64              `json:"pgScore" bson:"pgScore"`
	PgYear           int                  `json:"pgPassingyear" bson:"pgPassingyear"`
	UgIn             string               `json:"ugIn" bson:"ugIn"`
	UgCollege        string               `json:"ugCollege" bson:"ugCollege"`
	UgScore          float64              `json:"ugScore" bson:"ugScore"`
	UgYear           int                  `json:"ugPassingyear" bson:"ugPassingyear"`
	VideoResume      *string              `json:"videoResume" bson:"videoResume"`
	SemesterOne      *float64             `json:"semesterOne" bson:"semesterOne"`
	SemesterTwo      *float64             `json:"semesterTwo" bson:"semesterTwo"`
	SemesterThree    *float64             `json:"semesterThree" bson:"semesterThree"`
	SemesterFour     *float64             `json:"semesterFour" bson:"semesterFour"`
	SemesterFive     *float64             `json:"semesterFive" bson:"semesterFive"`
	SemesterSix      *float64             `json:"semesterSix" bson:"semesterSix"`
	SemesterSeven    *float64             `json:"semesterSeven" bson:"semesterSeven"`
	SemesterEight    *float64             `json:"semesterEight" bson:"semesterEight"`
	SummerOne        *float64             `json:"summerOne" bson:"summerOne"`
	SummerTwo        *float64             `json:"summerTwo" bson:"summerTwo"`
	SummerThree      *float64             `json:"summerThree" bson:"summerThree"`
	SummerFour       *float64             `json:"summerFour" bson:"summerFour"`
	SummerFive       *float64             `json:"summerFive" bson:"summerFive"`
	UpdatedAt        primitive.DateTime   `json:"updatedAt" bson:"updatedAt"`
}

type XXIIYearType1 struct {
	XiiYear string `json:"xiiYear" bson:"xiiYear"`
	XYear   string `json:"xYear" bson:"xYear"`
}

type XXIIYearType2 struct {
	XiiYear int `json:"xiiYear" bson:"xiiYear"`
	XYear   int `json:"xYear" bson:"xYear"`
}

type EuducationGapJeeRankIntegers struct {
	JeeRank      int `json:"jeeRank" bson:"jeeRank"`
	EducationGap int `json:"educationGap" bson:"educationGap"`
}

type EuducationGapJeeRankStrings struct {
	JeeRank      string `json:"jeeRank" bson:"jeeRank"`
	EducationGap string `json:"educationGap" bson:"educationGap"`
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
	SummerOne        float64            `json:"summerOne" bson:"summerOne"`
	SummerTwo        float64            `json:"summerTwo" bson:"summerTwo"`
	SummerThree      float64            `json:"summerThree" bson:"summerThree"`
	SummerFour       float64            `json:"summerFour" bson:"summerFour"`
	SummerFive       float64            `json:"summerFive" bson:"summerFive"`
	Groups           []Group            `json:"groups" bson:"groups"`
	UpdatedAt        primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
	CompaniesAlloted []string           `json:"companiesAlloted" bson:"companiesAlloted"`
}
