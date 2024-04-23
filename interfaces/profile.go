package interfaces

type GenericField struct {
	DataType    string    `json:"dataType"`
	DataChoices *[]string `json:"dataChoices,omitempty"`
	IsVerified  *bool     `json:"isVerified,omitempty"`
	Value       any       `json:"value"`
	IsLocked    bool      `json:"isEditable"`
	IsRequired  bool      `json:"isRequired"`
	IsNull      bool      `json:"isNull"`
	IsHidden    bool      `json:"isHidden"`
}

type ProfilePersonal struct {
	FirstName        GenericField `json:"firstName,omitempty"`
	MiddleName       GenericField `json:"middleName,omitempty"`
	LastName         GenericField `json:"lastName,omitempty"`
	Gender           GenericField `json:"gender,omitempty"`
	DOB              GenericField `json:"dob,omitempty"`
	PermanentAddress GenericField `json:"permanentAddress,omitempty"`
	PresentAddress   GenericField `json:"presentAddress,omitempty"`
	PersonalEmail    GenericField `json:"personalEmail,omitempty"`
	Mobile           GenericField `json:"mobile,omitempty"`
	Category         GenericField `json:"category,omitempty"`
	IsPWD            GenericField `json:"isPWD,omitempty"`
	IsEWS            GenericField `json:"isEWS,omitempty"`
	FatherName       GenericField `json:"fatherName,omitempty"`
	FatherOccupation GenericField `json:"fatherOccupation,omitempty"`
	MotherName       GenericField `json:"motherName,omitempty"`
	MotherOccupation GenericField `json:"motherOccupation,omitempty"`
	MotherTongue     GenericField `json:"motherTongue,omitempty"`
}

type ProfileSocials struct {
	LinkedIn       GenericField `json:"linkedIn,omitempty"`
	Github         GenericField `json:"github,omitempty"`
	Kaggle         GenericField `json:"kaggle,omitempty"`
	MicrosoftTeams GenericField `json:"microsoftTeams,omitempty"`
	Skype          GenericField `json:"skype,omitempty"`
	GoogleScholar  GenericField `json:"googleScholar,omitempty"`
	Codeforces     GenericField `json:"codeforces,omitempty"`
	CodeChef       GenericField `json:"codechef,omitempty"`
	Leetcode       GenericField `json:"leetcode,omitempty"`
}

type GenericRank struct {
	Rank     GenericField `json:"rank,omitempty"`
	Category GenericField `json:"category,omitempty"`
	IsEWS    GenericField `json:"isEWS,omitempty"`
	IsPWD    GenericField `json:"isPWD,omitempty"`
}

type ProfileInstitute struct {
	RollNumber     GenericField `json:"rollNo,omitempty"`
	InstituteEmail GenericField `json:"email,omitempty"`
	Batch          GenericField `json:"batch,omitempty"`
	Department     GenericField `json:"department,omitempty"`
	Course         GenericField `json:"course,omitempty"`
	Specialisation GenericField `json:"specialisation,omitempty"`
	Honours        GenericField `json:"honours,omitempty"`
	ThesisEndDate  GenericField `json:"thesisEndDate,omitempty"`
	EducationGap   GenericField `json:"educationGap,omitempty"`
}

type ProfileDetails struct {
	PersonalProfile  ProfilePersonal  `json:"personal_profile,omitempty"`
	SocialProfile    ProfileSocials   `json:"social_profile,omitempty"`
	InstituteProfile ProfileInstitute `json:"institute_profile,omitempty"`
}

type ProfilePastEducation struct {
	Certification GenericField `json:"certification,omitempty"`
	Institute     GenericField `json:"institute,omitempty"`
	Year          GenericField `json:"year,omitempty"`
	Score         GenericField `json:"score,omitempty"`
}

type ProfilePastAcademics struct {
	ClassX        ProfilePastEducation `json:"class_x,omitempty"`
	ClassXII      ProfilePastEducation `json:"class_xii,omitempty"`
	Undergraduate ProfilePastEducation `json:"undergraduate,omitempty"`
	Postgraduate  ProfilePastEducation `json:"postgraduate,omitempty"`
	JeeRank       GenericRank          `json:"jeeRank,omitempty"`
	GateRank      GenericRank          `json:"gateRank,omitempty"`
}

type ProfileSemesterSPI struct {
	One   GenericField `json:"one,omitempty"`
	Two   GenericField `json:"two,omitempty"`
	Three GenericField `json:"three,omitempty"`
	Four  GenericField `json:"four,omitempty"`
	Five  GenericField `json:"five,omitempty"`
	Six   GenericField `json:"six,omitempty"`
	Seven GenericField `json:"seven,omitempty"`
	Eight GenericField `json:"eight,omitempty"`
	Nine  GenericField `json:"nine,omitempty"`
	Ten   GenericField `json:"ten,omitempty"`
}

type ProfileSummerTermSPI struct {
	One   GenericField `json:"one,omitempty"`
	Two   GenericField `json:"two,omitempty"`
	Three GenericField `json:"three,omitempty"`
	Four  GenericField `json:"four,omitempty"`
	Five  GenericField `json:"five,omitempty"`
}

type ProfileCurrentAcademics struct {
	SemesterSPI   ProfileSemesterSPI   `json:"semester_spi,omitempty"`
	SummerTermSPI ProfileSummerTermSPI `json:"summer_term_spi,omitempty"`
}

type StudentProfile struct {
	Profile          ProfileDetails          `json:"profile,omitempty"`
	PastAcademics    ProfilePastAcademics    `json:"past_academics,omitempty"`
	CurrentAcademics ProfileCurrentAcademics `json:"current_academics,omitempty"`
}
