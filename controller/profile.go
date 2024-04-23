package controller

import (
	"fmt"
	"reflect"

	"github.com/FrosTiK-SD/auth/interfaces"
	"github.com/FrosTiK-SD/auth/model"
	studentModel "github.com/FrosTiK-SD/models/student"
	"github.com/modern-go/reflect2"
)

func AssignReservationCategory(category *interfaces.GenericField, isEWS *interfaces.GenericField, isPWD *interfaces.GenericField, rc *studentModel.ReservationCategory) {
	if reflect2.IsNil(rc) {
		category.IsNull = true
		isEWS.IsNull = true
		isPWD.IsNull = true
	} else {
		category.Value = rc.Category
		isEWS.Value = rc.IsEWS
		isPWD.Value = rc.IsPWD
	}
}

func AssignSocialProfile(field *interfaces.GenericField, social *studentModel.SocialProfile) {
	field.IsNull = reflect2.IsNil(field)
	field.DataType = "Social"

	if !field.IsNull {
		field.Value = social.URL + "|" + social.Username
	}
}

func AssignNilPossibleValue(field *interfaces.GenericField, value any) {
	field.Value = value
	field.IsNull = reflect2.IsNil(value)
	field.DataType = fmt.Sprintf("%v", reflect.TypeOf(value))
}

func AssignNotNilValue(field *interfaces.GenericField, value any) {
	field.Value = value
	field.DataType = fmt.Sprintf("%v", reflect.TypeOf(value))
}

func MapProfilePersonal(profile *interfaces.ProfilePersonal, student *model.StudentPopulated) {
	AssignNotNilValue(&profile.FirstName, student.FirstName)
	AssignNilPossibleValue(&profile.MiddleName, student.MiddleName)
	AssignNilPossibleValue(&profile.LastName, student.LastName)

	AssignNilPossibleValue(&profile.Gender, student.Gender)
	AssignNilPossibleValue(&profile.DOB, student.DOB)
	AssignNotNilValue(&profile.PermanentAddress, student.PermanentAddress)
	AssignNotNilValue(&profile.PresentAddress, student.PresentAddress)
	AssignNotNilValue(&profile.PersonalEmail, student.PersonalEmail)
	AssignNilPossibleValue(&profile.Mobile, student.Mobile)
	AssignReservationCategory(&profile.Category, &profile.IsEWS, &profile.IsPWD, student.Category)
	AssignNotNilValue(&profile.MotherTongue, student.MotherTongue)

	// required
	profile.FirstName.IsRequired = true
	profile.DOB.IsRequired = true
	profile.PermanentAddress.IsRequired = true
	profile.PersonalEmail.IsRequired = true
	profile.Mobile.IsRequired = true
}

func MapProfileCurrentAcademics(profile *interfaces.ProfileCurrentAcademics, academics *studentModel.Academics) {
	AssignNilPossibleValue(&profile.SemesterSPI.One, academics.SemesterSPI.One)
	AssignNilPossibleValue(&profile.SemesterSPI.Two, academics.SemesterSPI.Two)
	AssignNilPossibleValue(&profile.SemesterSPI.Three, academics.SemesterSPI.Three)
	AssignNilPossibleValue(&profile.SemesterSPI.Four, academics.SemesterSPI.Four)
	AssignNilPossibleValue(&profile.SemesterSPI.Five, academics.SemesterSPI.Five)
	AssignNilPossibleValue(&profile.SemesterSPI.Six, academics.SemesterSPI.Six)
	AssignNilPossibleValue(&profile.SemesterSPI.Seven, academics.SemesterSPI.Seven)
	AssignNilPossibleValue(&profile.SemesterSPI.Eight, academics.SemesterSPI.Eight)

	AssignNilPossibleValue(&profile.SummerTermSPI.One, academics.SummerTermSPI.One)
	AssignNilPossibleValue(&profile.SummerTermSPI.Two, academics.SummerTermSPI.Two)
	AssignNilPossibleValue(&profile.SummerTermSPI.Three, academics.SummerTermSPI.Three)
	AssignNilPossibleValue(&profile.SummerTermSPI.Four, academics.SummerTermSPI.Four)
	AssignNilPossibleValue(&profile.SummerTermSPI.Five, academics.SummerTermSPI.Five)
}

func MapProfileSocials(profile *interfaces.ProfileSocials, socials *studentModel.SocialProfiles) {
	
}

func MapStudentToStudentProfile(student *model.StudentPopulated) interfaces.StudentProfile {
	var profile interfaces.StudentProfile
	MapProfilePersonal(&profile.Profile.PersonalProfile, student)
	MapProfileCurrentAcademics(&profile.CurrentAcademics, &student.Academics)
	return profile
}
