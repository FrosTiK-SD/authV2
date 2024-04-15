package interfaces

type GenericField struct {
	DataType    string      `json:"dataType"`
	DataChoices *[]string   `json:"dataChoices,omitempty"`
	IsVerified  *bool       `json:"isVerified,omitempty"`
	Value       interface{} `json:"value"`
	IsEditable  bool        `json:"isEditable"`
	IsRequired  bool        `json:"isRequired"`
	IsNull      bool        `json:"isNull"`
}

type ProfilePersonal struct {
	FirstName        GenericField `json:"first_name,omitempty"`
	MiddleName       GenericField `json:"middle_name,omitempty"`
	LastName         GenericField `json:"last_name,omitempty"`
	Gender           GenericField `json:"gender,omitempty"`
	PermanentAddress GenericField `json:"permanent_address,omitempty"`
	PresentAddress   GenericField `json:"present_address,omitempty"`
	Mobile           GenericField `json:"mobile,omitempty"`
	Category         GenericField `json:"category,omitempty"`
	IsPWD            GenericField `json:"is_pwd,omitempty"`
	IsEWS            GenericField `json:"is_ews,omitempty"`
	MotherTongue     GenericField `json:"mother_tongue,omitempty"`
}

type ProfileSocials struct {
	LinkedIn GenericField `json:"linked_in,omitempty"`
	Github   GenericField `json:"github,omitempty"`
}

type ProfileInstitute struct {
	RollNumber     GenericField `json:"roll_number,omitempty"`
	InstituteEmail GenericField `json:"institute_email,omitempty"`
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
