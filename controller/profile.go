package controller

import (
	"fmt"
	"reflect"

	"strconv"

	"github.com/FrosTiK-SD/auth/interfaces"
	"github.com/FrosTiK-SD/auth/model"
	"github.com/FrosTiK-SD/models/student"
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
	if social != nil {
		field.IsNull = reflect2.IsNil(field)
		field.DataType = "Social"

		if !field.IsNull {
			field.Value = social.URL + "|" + social.Username
		}
	}
	if social == nil {
		field.IsNull = true
		field.DataType = "Social"
		field.Value = "null"
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

func AssignPastAcademics(field *interfaces.ProfilePastEducation, education *student.EducationDetails) {
	if education == nil {
       AssignNilPossibleValue(&field.Certification,education)
       AssignNilPossibleValue(&field.Institute,education)
	   AssignNilPossibleValue(&field.Year,education)
	   AssignNilPossibleValue(&field.Score,education)
	}
	if education != nil {
		AssignNotNilValue(&field.Certification,education.Certification)
        AssignNotNilValue(&field.Institute,education.Institute)
        AssignNotNilValue(&field.Year,education.Year)
        AssignNotNilValue(&field.Score,education.Score)
	}
}

func AssignRankValue(rankField *interfaces.GenericRank, rankDetails *student.RankDetails){
	if rankDetails == nil {
        AssignNilPossibleValue(&rankField.Rank,rankDetails)
    }
    if rankDetails != nil {
        AssignNotNilValue(&rankField.Rank,rankDetails.Rank)
		AssignReservationCategory(&rankField.Rank,&rankField.IsEWS,&rankField.IsPWD,&rankDetails.RankCategory)
    }
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
	AssignNotNilValue(&profile.FatherName, student.ParentsDetails.FatherName)
	AssignNotNilValue(&profile.MotherName, student.ParentsDetails.MotherName)
	AssignNotNilValue(&profile.FatherOccupation, student.ParentsDetails.FatherOccupation)
	AssignNotNilValue(&profile.MotherOccupation, student.ParentsDetails.MotherOccupation)
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

func AssignBatch(profile *interfaces.GenericField, institute *studentModel.Student) {
	profile.IsNull = reflect2.IsNil(institute.Batch)
	profile.DataType = "string"
	if !profile.IsNull {
		profile.Value = strconv.Itoa(institute.Batch.StartYear) + "-" + strconv.Itoa(institute.Batch.EndYear)
	}
}

func MapProfileSocials(profile *interfaces.ProfileSocials, socials *studentModel.SocialProfiles) {
	AssignSocialProfile(&profile.LinkedIn, socials.LinkedIn)
	AssignSocialProfile(&profile.Github, socials.Github)
	AssignSocialProfile(&profile.CodeChef, socials.CodeChef)
	AssignSocialProfile(&profile.Codeforces, socials.Codeforces)
	AssignSocialProfile(&profile.Leetcode, socials.LeetCode)
	AssignSocialProfile(&profile.GoogleScholar, socials.GoogleScholar)
	AssignSocialProfile(&profile.MicrosoftTeams, socials.MicrosoftTeams)
	AssignSocialProfile(&profile.Kaggle, socials.Kaggle)
	AssignSocialProfile(&profile.Skype, socials.Skype)
}

func MapProfileInstitute(profile *interfaces.ProfileInstitute, institute *studentModel.Student) {
	AssignBatch(&profile.Batch, institute)
	AssignNotNilValue(&profile.RollNumber, institute.RollNo)
	AssignNotNilValue(&profile.InstituteEmail, institute.InstituteEmail)
	AssignNotNilValue(&profile.Department, institute.Department)
	AssignNilPossibleValue(&profile.EducationGap, institute.Academics.EducationGap)
	AssignNotNilValue(&profile.Course, institute.Course)
	AssignNilPossibleValue(&profile.Specialisation, institute.Specialisation)
	AssignNilPossibleValue(&profile.Honours, institute.Academics.Honours)
	AssignNilPossibleValue(&profile.ThesisEndDate, institute.Academics.ThesisEndDate)
}

func MapPastAcademics(profile *interfaces.ProfilePastAcademics, institute *studentModel.Academics) {
	AssignPastAcademics(&profile.ClassX, institute.XthClass)
	AssignPastAcademics(&profile.ClassXII, institute.XIIthClass)
	AssignPastAcademics(&profile.Undergraduate, institute.UnderGraduate)
	AssignPastAcademics(&profile.Postgraduate, institute.PostGraduate)
}

func MapRanks(profile *interfaces.ProfileInstitute,institute *studentModel.Academics){
	AssignRankValue(&profile.JeeRank,institute.JEERank)
	AssignRankValue(&profile.GateRank,institute.GATERank)
}

func MapStudentToStudentProfile(student *model.StudentPopulated) interfaces.StudentProfile {
	var profile interfaces.StudentProfile
	MapProfilePersonal(&profile.Profile.PersonalProfile, student)
	MapProfileCurrentAcademics(&profile.CurrentAcademics, &student.Academics)
	MapProfileSocials(&profile.Profile.SocialProfile, &student.SocialProfiles)
	MapProfileInstitute(&profile.Profile.InstituteProfile, &student.Student)
	MapPastAcademics(&profile.PastAcademics, &student.Academics)
	MapRanks(&profile.Profile.InstituteProfile, &student.Academics)
	return profile
}
