package model

import (
	constant "github.com/FrosTiK-SD/models/constant"
	misc "github.com/FrosTiK-SD/models/misc"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationCategory struct {
	Category string `json:"category" bson:"category"`
	IsPWD    bool   `json:"isPWD" bson:"isPWD"`
	IsEWS    bool   `json:"isEWS" bson:"isEWS"`
}

type RankDetails struct {
	Rank         int                 `bson:"rank" json:"rank"`
	RankCategory ReservationCategory `bson:"rankCategory" json:"rankCategory"`
}

type SemesterSPI struct {
	One   float64 `json:"One,omitempty" bson:"One,omitempty"`
	Two   float64 `json:"Two,omitempty" bson:"Two,omitempty"`
	Three float64 `json:"Three,omitempty" bson:"Three,omitempty"`
	Four  float64 `json:"Four,omitempty" bson:"Four,omitempty"`
	Five  float64 `json:"Five,omitempty" bson:"Five,omitempty"`
	Six   float64 `json:"Six,omitempty" bson:"Six,omitempty"`
	Seven float64 `json:"Seven,omitempty" bson:"Seven,omitempty"`
	Eight float64 `json:"Eight,omitempty" bson:"Eight,omitempty"`
	Nine  float64 `json:"Nine,omitempty" bson:"Nine,omitempty"`
	Ten   float64 `json:"Ten,omitempty" bson:"Ten,omitempty"`
}

type SummerTermSPI struct {
	One   float64 `json:"One,omitempty" bson:"One,omitempty"`
	Two   float64 `json:"Two,omitempty" bson:"Two,omitempty"`
	Three float64 `json:"Three,omitempty" bson:"Three,omitempty"`
	Four  float64 `json:"Four,omitempty" bson:"Four,omitempty"`
	Five  float64 `json:"Five,omitempty" bson:"Five,omitempty"`
}

type EducationDetails struct {
	Certification string  `bson:"certification" json:"certification"`
	Institute     string  `bson:"institute" json:"institute"`
	Year          int     `bson:"year" json:"year"`
	Score         float64 `bson:"score" json:"score"`
}

type Academics struct {
	JEERank           *RankDetails        `json:"jeeRank" bson:"jeeRank"`
	GATERank          *RankDetails        `json:"gateRank,omitempty" bson:"gateRank,omitempty"`
	XthClass          *EducationDetails   `json:"xClass" bson:"xClass"`
	XIIthClass        *EducationDetails   `json:"xiiClass" bson:"xiiClass"`
	UnderGraduate     *EducationDetails   `json:"underGraduate,omitempty" bson:"underGraduate,omitempty"`
	Honours           *string             `json:"honours,omitempty" bson:"honours,omitempty"`
	PostGraduate      *EducationDetails   `json:"postGraduate,omitempty" bson:"postGraduate,omitempty"`
	ThesisEndDate     *primitive.DateTime `json:"thesisEndDate,omitempty" bson:"thesisEndDate,omitempty"`
	EducationGap      int                 `json:"educationGap" bson:"educationGap"`
	SemesterDetails   SemesterSPI         `json:"semesterSPI" bson:"semesterSPI"`
	SummerTermDetails SummerTermSPI       `json:"summerTermSPI" bson:"summerTermSPI"`
	CurrentCGPA       float64             `json:"currentCGPA" bson:"currentCGPA"`
	ActiveBacklogs    int                 `json:"activeBacklogs" bson:"activeBacklogs"`
	TotalBacklogs     int                 `json:"totalBacklogs" bson:"totalBacklogs"`

	Verification misc.Verification `json:"verification" bson:"verification"`
}

type SocialProfile struct {
	URL          string            `bson:"url" json:"url"`
	Username     string            `bson:"username" json:"username"`
	Verification misc.Verification `bson:"verification" json:"verification"`
}

type SocialProfiles struct {
	LinkedIn       *SocialProfile `json:"linkedIn" bson:"linkedIn"`
	Github         *SocialProfile `json:"github,omitempty" bson:"github,omitempty"`
	MicrosoftTeams *SocialProfile `json:"microsoftTeams,omitempty" bson:"microsoftTeams,omitempty"`
	Skype          *SocialProfile `json:"skype,omitempty" bson:"skype,omitempty"`
	GoogleScholar  *SocialProfile `json:"googleScholar,omitempty" bson:"googleScholar,omitempty"`

	Codeforces *SocialProfile `json:"codeforces,omitempty" bson:"codeforces,omitempty"`
	CodeChef   *SocialProfile `json:"codechef,omitempty" bson:"codechef,omitempty"`
	LeetCode   *SocialProfile `json:"leetcode,omitempty" bson:"leetcode,omitempty"`
	Kaggle     *SocialProfile `json:"kaggle,omitempty" bson:"kaggle,omitempty"`
}

type Batch struct {
	StartYear int `json:"startYear" bson:"startYear"`
	EndYear   int `json:"endYear" bson:"endYear"`
}

type ParentsDetails struct {
	FatherName       string `json:"fatherName" bson:"fatherName"`
	FatherOccupation string `json:"fatherOccupation,omitempty" bson:"fatherOccupation,omitempty"`
	MotherName       string `json:"motherName" bson:"motherName"`
	MotherOccupation string `json:"motherOccupation,omitempty" bson:"motherOccupation,omitempty"`
}

type Extras struct {
	VideoResume  string            `bson:"videoResume" json:"videoResume"`
	Verification misc.Verification `bson:"verification" json:"verification"`
}

type WorkExperience struct {
	StartDate    primitive.DateTime `bson:"startDate" json:"startDate"`
	EndDate      primitive.DateTime `bson:"endDate" json:"endDate"`
	Organisation string             `bson:"organisation" json:"organisation"`
	Location     string             `bson:"location" json:"location"`
	Position     string             `bson:"position" json:"position"`
	Details      string             `bson:"details" json:"details"`
	Verification misc.Verification  `bson:"verification" json:"verification"`
}

type Student struct {
	Id               primitive.ObjectID   `json:"_id" bson:"_id"`
	Groups           []primitive.ObjectID `json:"groups" bson:"groups"`
	CompaniesAlloted []string             `json:"companiesAlloted" bson:"companiesAlloted"`

	Batch          Batch           `json:"batch" bson:"batch"`
	RollNo         int             `json:"rollNo" bson:"rollNo"`
	InstituteEmail string          `json:"email" bson:"email"`
	Department     string          `json:"department" bson:"department"`
	Course         constant.Course `json:"course" bson:"course"`
	Specialisation *string         `json:"specialisation,omitempty" bson:"specialisation,omitempty"`

	FirstName  string `json:"firstName" bson:"firstName"`
	MiddleName string `json:"middleName,omitempty" bson:"middleName,omitempty"`
	LastName   string `json:"lastName" bson:"lastName"`

	ProfilePicture   *misc.Attachment     `json:"profilePicture,omitempty" bson:"profilePicture,omitempty"`
	Gender           constant.Gender      `json:"gender" bson:"gender"`
	DOB              *primitive.DateTime  `json:"dob" bson:"dob"`
	PermanentAddress string               `json:"permanentAddress" bson:"permanentAddress"`
	PresentAddress   string               `json:"presentAddress" bson:"presentAddress"`
	PersonalEmail    string               `json:"personalEmail" bson:"personalEmail"`
	Mobile           string               `json:"mobile" bson:"mobile"`
	Category         *ReservationCategory `json:"category" bson:"category"`
	MotherTongue     string               `json:"motherTongue" bson:"motherTongue"`
	ParentsDetails   *ParentsDetails      `json:"parentsDetails" bson:"parentsDetails"`

	Academics      Academics        `json:"academics" bson:"academics"`
	WorkExperience []WorkExperience `json:"workExperience" bson:"workExperience"`
	SocialProfiles SocialProfiles   `json:"socialProfiles" bson:"socialProfiles"`

	// metadata
	StructVersion int                    `json:"version,omitempty" bson:"version,omitempty"`
	UpdatedAt     primitive.DateTime     `json:"updatedAt" bson:"updatedAt"`
	CreatedAt     primitive.DateTime     `json:"createdAt" bson:"createdAt"`
	RawKeyStore   map[string]interface{} `json:"raw_key_store" bson:"raw_key_store"`
	DataErrors    []string               `json:"dataErrors" bson:"dataErrors"`
}

// at some point I expect to use * pointers to omit fields which are not necessary
// https://willnorris.com/2014/go-rest-apis-and-pointers/
// https://stackoverflow.com/questions/47158987/how-to-update-mongodb-fields-with-omitempty-flag-in-golang-structure
