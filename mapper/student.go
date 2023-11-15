package mapper

import (
	"fmt"

	"frostik.com/auth/model"
)

func TransformStudentToStudentPopulated(student model.Student) model.StudentPopulated {
	fmt.Println(student.CompaniesAlloted)
	return model.StudentPopulated{
		ID:               student.ID,
		Batch:            student.Batch,
		RollNo:           student.RollNo,
		FirstName:        student.FirstName,
		LastName:         student.LastName,
		Department:       student.Department,
		Course:           student.Course,
		Email:            student.Email,
		PersonalEmail:    student.PersonalEmail,
		LinkedIn:         student.LinkedIn,
		Github:           student.Github,
		MicrosoftTeams:   student.MicrosoftTeams,
		Mobile:           student.Mobile,
		Gender:           student.Gender,
		Dob:              student.Dob,
		PermanentAddress: student.PermanentAddress,
		PresentAddress:   student.PresentAddress,
		Category:         student.Category,
		FatherName:       student.FatherName,
		FatherOccupation: student.FatherOccupation,
		MotherName:       student.MotherName,
		MotherOccupation: student.MotherOccupation,
		MotherTongue:     student.MotherTongue,
		EducationGap:     student.EducationGap,
		JeeRank:          student.JeeRank,
		Cgpa:             student.Cgpa,
		ActiveBacklogs:   student.ActiveBacklogs,
		TotalBacklogs:    student.TotalBacklogs,
		XBoard:           student.XBoard,
		XYear:            student.XYear,
		XPercentage:      student.XPercentage,
		XInstitute:       student.XInstitute,
		XiiBoard:         student.XiiBoard,
		XiiYear:          student.XiiYear,
		XiiPercentage:    student.XiiPercentage,
		XiiInstitute:     student.XiiInstitute,
		SemesterOne:      student.SemesterOne,
		SemesterTwo:      student.SemesterTwo,
		SemesterThree:    student.SemesterThree,
		SemesterFour:     student.SemesterFour,
		SemesterFive:     student.SemesterFive,
		SemesterSix:      student.SemesterSix,
		Groups:           []model.Group{},
		CompaniesAlloted: student.CompaniesAlloted,
		UpdatedAt:        student.UpdatedAt,
	}
}
