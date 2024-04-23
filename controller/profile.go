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

func AssignNilPossibleValue(field *interfaces.GenericField, value any) {
	field.Value = nil
	field.IsNull = reflect2.IsNil(value)
	field.DataType = fmt.Sprintf("%v", reflect.TypeOf(value))
}

func AssignNotNilValue(field *interfaces.GenericField, value any) {
	field.Value = value
	field.DataType = fmt.Sprintf("%v", reflect.TypeOf(value))
}

func MapProfileDetails(profile *interfaces.StudentProfile, student *model.StudentPopulated) {
	AssignNotNilValue(&profile.Profile.PersonalProfile.FirstName, student.FirstName)
	AssignNotNilValue(&profile.Profile.InstituteProfile.RollNumber,student.RollNo)
	AssignNotNilValue(&profile.Profile.InstituteProfile.InstituteEmail,student.InstituteEmail)
	AssignNotNilValue(&profile.Profile.PersonalProfile.PermanentAddress,student.PermanentAddress)
	AssignNilPossibleValue(&profile.Profile.PersonalProfile.MiddleName, student.MiddleName)
	AssignNilPossibleValue(&profile.Profile.PersonalProfile.LastName, student.LastName)
	AssignNilPossibleValue(&profile.Profile.PersonalProfile.Gender, student.Gender)
	AssignNilPossibleValue(&profile.Profile.PersonalProfile.DOB, student.DOB)
	AssignNotNilValue(&profile.Profile.PersonalProfile.PresentAddress, student.PresentAddress)
	AssignNotNilValue(&profile.Profile.PersonalProfile.PersonalEmail, student.PersonalEmail)
	AssignNilPossibleValue(&profile.Profile.PersonalProfile.Mobile, student.Mobile)
	AssignReservationCategory(&profile.Profile.PersonalProfile.Category, &profile.Profile.PersonalProfile.IsEWS, &profile.Profile.PersonalProfile.IsPWD, student.Category)
	AssignNotNilValue(&profile.Profile.PersonalProfile.MotherTongue, student.MotherTongue)
	AssignNilPossibleValue(&profile.Profile.InstituteProfile.Specialisation, student.Specialisation)
	AssignNilPossibleValue(&profile.Profile.InstituteProfile.Honours,student.Academics.Honours)
	AssignNilPossibleValue(&profile.Profile.InstituteProfile.EducationGap,student.Academics.EducationGap)
	AssignNilPossibleValue(&profile.Profile.InstituteProfile.ThesisEndDate,student.Academics.ThesisEndDate)
    AssignNotNilValue(&profile.Profile.PersonalProfile.MotherName,student.ParentsDetails.MotherName)
	AssignNotNilValue(&profile.Profile.PersonalProfile.FatherName,student.ParentsDetails.FatherName)
	AssignNotNilValue(&profile.Profile.PersonalProfile.MotherOccupation,student.ParentsDetails.MotherOccupation)
	AssignNotNilValue(&profile.Profile.PersonalProfile.FatherOccupation,student.ParentsDetails.FatherOccupation)
	AssignNilPossibleValue(&profile.CurrentAcademics.SemesterSPI.Seven,student.Academics.SemesterSPI.Seven)
	AssignNilPossibleValue(&profile.CurrentAcademics.SemesterSPI.Eight,student.Academics.SemesterSPI.Eight)
	AssignNilPossibleValue(&profile.CurrentAcademics.SemesterSPI.Nine,student.Academics.SemesterSPI.Nine)
	AssignNilPossibleValue(&profile.CurrentAcademics.SemesterSPI.Ten,student.Academics.SemesterSPI.Ten)
	AssignNilPossibleValue(&profile.CurrentAcademics.SummerTermSPI.One,student.Academics.SummerTermSPI.One)
	AssignNilPossibleValue(&profile.CurrentAcademics.SummerTermSPI.Two,student.Academics.SummerTermSPI.Two)
	AssignNilPossibleValue(&profile.CurrentAcademics.SummerTermSPI.Three,student.Academics.SummerTermSPI.Three)
	AssignNilPossibleValue(&profile.CurrentAcademics.SummerTermSPI.Four,student.Academics.SummerTermSPI.Four)
	AssignNilPossibleValue(&profile.CurrentAcademics.SummerTermSPI.Five,student.Academics.SummerTermSPI.Five)

	// required
	profile.Profile.PersonalProfile.FirstName.IsRequired = true
	profile.Profile.PersonalProfile.DOB.IsRequired = true
	profile.Profile.PersonalProfile.PermanentAddress.IsRequired = true
	profile.Profile.PersonalProfile.PersonalEmail.IsRequired = true
	profile.Profile.PersonalProfile.Mobile.IsRequired = true
}

func MapStudentToStudentProfile(student *model.StudentPopulated) interfaces.StudentProfile {
	var profile interfaces.StudentProfile
	MapProfileDetails(&profile, student)
	return profile
}