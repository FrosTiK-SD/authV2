package mapper

import (
	"github.com/FrosTiK-SD/auth/model"
)

func TransformStudentToStudentPopulated(student model.OldStudent) model.StudentPopulated {
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
		XBoard:           student.XBoard,
		XPercentage:      student.XPercentage,
		XInstitute:       student.XInstitute,
		XiiBoard:         student.XiiBoard,
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
