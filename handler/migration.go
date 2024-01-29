package handler

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/FrosTiK-SD/auth/model"
	Constant "github.com/FrosTiK-SD/models/constant"
	Student "github.com/FrosTiK-SD/models/student"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetDOBFromString(input string) (primitive.DateTime, error) {
	t, err := time.Parse("02-01-2006", input)
	if err != nil {
		return 0, err
	}
	return primitive.NewDateTimeFromTime(t), nil
}

func GetCategoryFromString(input string) Student.ReservationCategory {
	category := Student.ReservationCategory{
		Category: "GEN",
		IsPWD:    false,
		IsEWS:    false,
	}
	for _, c := range [4]string{"GEN", "SC", "ST", "OBC-NCL"} {
		re := regexp.MustCompile("(?i)" + c)

		if re.MatchString(input) {
			category.Category = c
			break
		}
	}

	if regexp.MustCompile("(?i)pwd").MatchString(input) {
		category.IsPWD = true
	}

	if regexp.MustCompile("(?i)ews").MatchString(input) {
		category.IsEWS = true
	}

	return category
}

func GetRankFromString(input string, rc Student.ReservationCategory) (Student.RankDetails, error) {
	srd := Student.RankDetails{
		Rank:         -1,
		RankCategory: GetCategoryFromString(input),
	}

	num, err := strconv.Atoi(regexp.MustCompile(`\d+`).FindString(input))
	if err != nil {
		return srd, err
	}
	srd.Rank = num

	if regexp.MustCompile("(?i)category").MatchString(input) {
		srd.RankCategory = rc
	}

	return srd, nil

}

func (h *Handler) MigrateStudentDataToV2FormatType1(ctx *gin.Context) {
	studentCollection := h.MongikClient.MongoClient.Database(Constant.DB).Collection(Constant.StudentCollection)
	cursor, errFind := studentCollection.Find(ctx, bson.D{{Key: "version", Value: bson.D{{Key: "$exists", Value: false}}}})
	if errFind != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"count": 0, "error": errFind.Error(), "reason": "Find not successful"})
		return
	}

	var count int64 = 0

	for cursor.Next(ctx) {
		var oldStudent model.OldStudentType1

		if errDecode := cursor.Decode(&oldStudent); errDecode != nil {
			continue
		}

		if oldStudent.Email == "tpo@itbhu.ac.in" {
			continue
		}

		if oldStudent.CompaniesAlloted == nil {
			oldStudent.CompaniesAlloted = []string{}
		}

		errorArray := []string{"educationGap"}
		var EndYearOffset int
		var Course Constant.Course

		switch oldStudent.Course {
		case "idd":
			EndYearOffset = 5
			Course = Constant.IDD
		case "mtech":
			EndYearOffset = 2
			Course = Constant.MTECH
		case "phd":
			EndYearOffset = 6
			Course = Constant.PHD
		default:
			EndYearOffset = 4
			Course = Constant.BTECH
		}

		category := GetCategoryFromString(oldStudent.Category)

		jeeRank, errJeeRank := GetRankFromString(oldStudent.JeeRank, category)
		if errJeeRank != nil {
			fmt.Println(errJeeRank)
			errorArray = append(errorArray, "jeeRank")
			jeeRank.Rank = -1
		}

		xYear, xYearError := strconv.Atoi(oldStudent.XYear)
		if xYearError != nil {
			fmt.Println(xYearError)
			errorArray = append(errorArray, "xYear")
			xYear = -1
		}

		xiiYear, xiiYearError := strconv.Atoi(oldStudent.XiiYear)
		if xiiYearError != nil {
			fmt.Println(xiiYearError)
			errorArray = append(errorArray, "xiiYear")
			xiiYear = -1
		}

		dob, errDOB := GetDOBFromString(oldStudent.Dob)
		if errDOB != nil {
			fmt.Println(errDOB)
			errorArray = append(errorArray, "dob")
			dob = 0
		}

		newStudent := Student.Student{
			Id:               oldStudent.ID,
			Groups:           oldStudent.Groups,
			CompaniesAlloted: oldStudent.CompaniesAlloted,

			Batch: Student.Batch{
				StartYear: oldStudent.Batch,
				EndYear:   oldStudent.Batch + EndYearOffset,
			},
			RollNo:         oldStudent.RollNo,
			InstituteEmail: oldStudent.Email,
			Department:     oldStudent.Department,
			Course:         Course,

			FirstName: oldStudent.FirstName,
			LastName:  oldStudent.LastName,

			Gender:           Constant.Gender(strings.ToLower(oldStudent.Gender)),
			DOB:              dob,
			PermanentAddress: oldStudent.PermanentAddress,
			PresentAddress:   oldStudent.PresentAddress,
			PersonalEmail:    oldStudent.PersonalEmail,
			Mobile:           strconv.FormatInt(oldStudent.Mobile, 10),
			Category:         category,
			MotherTongue:     oldStudent.MotherTongue,
			ParentsDetails: Student.ParentsDetails{
				FatherName:       oldStudent.FatherName,
				FatherOccupation: oldStudent.FatherOccupation,
				MotherName:       oldStudent.MotherName,
				MotherOccupation: oldStudent.MotherOccupation,
			},

			Academics: Student.Academics{
				JEERank: jeeRank,
				XthClass: Student.EducationDetails{
					Certification: oldStudent.XBoard,
					Institute:     oldStudent.XInstitute,
					Year:          xYear,
					Score:         oldStudent.XPercentage,
				},
				XIIthClass: Student.EducationDetails{
					Certification: oldStudent.XiiBoard,
					Institute:     oldStudent.XiiInstitute,
					Year:          xiiYear,
					Score:         oldStudent.XiiPercentage,
				},
				EducationGap: -1,
				SemesterDetails: Student.SemesterSPI{
					One:   oldStudent.SemesterOne,
					Two:   oldStudent.SemesterTwo,
					Three: oldStudent.SemesterThree,
					Four:  oldStudent.SemesterFour,
					Five:  oldStudent.SemesterFive,
					Six:   oldStudent.SemesterSix,
				},
				SummerTermDetails: Student.SummerTermSPI{
					One:   oldStudent.SummerOne,
					Two:   oldStudent.SummerTwo,
					Three: oldStudent.SummerThree,
					Four:  oldStudent.SummerFour,
					Five:  oldStudent.SummerFive,
				},
				UnderGraduate: Student.EducationDetails{
					Certification: oldStudent.UgIn,
					Institute:     oldStudent.UgCollege,
					Year:          oldStudent.UgYear,
					Score:         oldStudent.UgScore,
				},
				PostGraduate: Student.EducationDetails{
					Certification: oldStudent.PgIn,
					Institute:     oldStudent.PgCollege,
					Year:          oldStudent.PgYear,
					Score:         oldStudent.PgScore,
				},
				CurrentCGPA:    oldStudent.Cgpa,
				ActiveBacklogs: oldStudent.ActiveBacklogs,
				TotalBacklogs:  oldStudent.TotalBacklogs,
			},
			WorkExperience: []Student.WorkExperience{},
			SocialProfiles: Student.SocialProfiles{
				LinkedIn: Student.SocialProfile{
					URL: oldStudent.LinkedIn,
				},
				MicrosoftTeams: Student.SocialProfile{
					URL: oldStudent.MicrosoftTeams,
				},
				Github: Student.SocialProfile{
					URL: oldStudent.Github,
				},
				Kaggle: Student.SocialProfile{
					URL: oldStudent.Kaggle,
				},
				Skype: Student.SocialProfile{
					URL: oldStudent.Skype,
				},
			},

			StructVersion: 2,
			CreatedAt:     primitive.NewDateTimeFromTime(time.Now()),
			UpdatedAt:     primitive.NewDateTimeFromTime(time.Now()),
			DataErrors:    errorArray,
		}

		filter := bson.D{{Key: "_id", Value: oldStudent.ID}}

		if _, errUpdate := studentCollection.ReplaceOne(ctx, filter, newStudent); errUpdate != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"count": count, "error": errUpdate.Error(), "id": oldStudent.ID})
			return
		}

		count = count + 1
	}

	ctx.JSON(200, gin.H{"count": count})
}

func (h *Handler) MigrateStudentDataToV2FormatType2(ctx *gin.Context) {
	studentCollection := h.MongikClient.MongoClient.Database(Constant.DB).Collection(Constant.StudentCollection)
	cursor, errFind := studentCollection.Find(ctx, bson.D{{Key: "version", Value: bson.D{{Key: "$exists", Value: false}}}})
	if errFind != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"count": 0, "error": errFind.Error(), "reason": "Find not successful"})
		return
	}

	var count int64 = 0

	for cursor.Next(ctx) {
		var oldStudent model.OldStudentType2

		if errDecode := cursor.Decode(&oldStudent); errDecode != nil {
			fmt.Println(errDecode.Error())
			continue
		}

		if oldStudent.Email == "tpo@itbhu.ac.in" {
			continue
		}

		if oldStudent.CompaniesAlloted == nil {
			oldStudent.CompaniesAlloted = []string{}
		}

		errorArray := []string{"educationGap"}
		var EndYearOffset int
		var Course Constant.Course

		switch oldStudent.Course {
		case "idd":
			EndYearOffset = 5
			Course = Constant.IDD
		case "mtech":
			EndYearOffset = 2
			Course = Constant.MTECH
		case "phd":
			EndYearOffset = 6
			Course = Constant.PHD
		default:
			EndYearOffset = 4
			Course = Constant.BTECH
		}

		category := GetCategoryFromString(oldStudent.Category)

		jeeRank, errJeeRank := GetRankFromString(oldStudent.JeeRank, category)
		if errJeeRank != nil {
			fmt.Println(errJeeRank)
			errorArray = append(errorArray, "jeeRank")
			jeeRank.Rank = -1
		}

		dob, errDOB := GetDOBFromString(oldStudent.Dob)
		if errDOB != nil {
			fmt.Println(errDOB)
			errorArray = append(errorArray, "dob")
			dob = 0
		}

		newStudent := Student.Student{
			Id:               oldStudent.ID,
			Groups:           oldStudent.Groups,
			CompaniesAlloted: oldStudent.CompaniesAlloted,

			Batch: Student.Batch{
				StartYear: oldStudent.Batch,
				EndYear:   oldStudent.Batch + EndYearOffset,
			},
			RollNo:         oldStudent.RollNo,
			InstituteEmail: oldStudent.Email,
			Department:     oldStudent.Department,
			Course:         Course,

			FirstName: oldStudent.FirstName,
			LastName:  oldStudent.LastName,

			Gender:           Constant.Gender(strings.ToLower(oldStudent.Gender)),
			DOB:              dob,
			PermanentAddress: oldStudent.PermanentAddress,
			PresentAddress:   oldStudent.PresentAddress,
			PersonalEmail:    oldStudent.PersonalEmail,
			Mobile:           strconv.FormatInt(oldStudent.Mobile, 10),
			Category:         category,
			MotherTongue:     oldStudent.MotherTongue,
			ParentsDetails: Student.ParentsDetails{
				FatherName:       oldStudent.FatherName,
				FatherOccupation: oldStudent.FatherOccupation,
				MotherName:       oldStudent.MotherName,
				MotherOccupation: oldStudent.MotherOccupation,
			},

			Academics: Student.Academics{
				JEERank: jeeRank,
				XthClass: Student.EducationDetails{
					Certification: oldStudent.XBoard,
					Institute:     oldStudent.XInstitute,
					Year:          oldStudent.XYear,
					Score:         oldStudent.XPercentage,
				},
				XIIthClass: Student.EducationDetails{
					Certification: oldStudent.XiiBoard,
					Institute:     oldStudent.XiiInstitute,
					Year:          oldStudent.XiiYear,
					Score:         oldStudent.XiiPercentage,
				},
				EducationGap: -1,
				SemesterDetails: Student.SemesterSPI{
					One:   oldStudent.SemesterOne,
					Two:   oldStudent.SemesterTwo,
					Three: oldStudent.SemesterThree,
					Four:  oldStudent.SemesterFour,
					Five:  oldStudent.SemesterFive,
					Six:   oldStudent.SemesterSix,
				},
				SummerTermDetails: Student.SummerTermSPI{
					One:   oldStudent.SummerOne,
					Two:   oldStudent.SummerTwo,
					Three: oldStudent.SummerThree,
					Four:  oldStudent.SummerFour,
					Five:  oldStudent.SummerFive,
				},
				UnderGraduate: Student.EducationDetails{
					Certification: oldStudent.UgIn,
					Institute:     oldStudent.UgCollege,
					Year:          oldStudent.UgYear,
					Score:         oldStudent.UgScore,
				},
				PostGraduate: Student.EducationDetails{
					Certification: oldStudent.PgIn,
					Institute:     oldStudent.PgCollege,
					Year:          oldStudent.PgYear,
					Score:         oldStudent.PgScore,
				},
				CurrentCGPA:    oldStudent.Cgpa,
				ActiveBacklogs: oldStudent.ActiveBacklogs,
				TotalBacklogs:  oldStudent.TotalBacklogs,
			},
			WorkExperience: []Student.WorkExperience{},
			SocialProfiles: Student.SocialProfiles{
				LinkedIn: Student.SocialProfile{
					URL: oldStudent.LinkedIn,
				},
				MicrosoftTeams: Student.SocialProfile{
					URL: oldStudent.MicrosoftTeams,
				},
				Github: Student.SocialProfile{
					URL: oldStudent.Github,
				},
				Kaggle: Student.SocialProfile{
					URL: oldStudent.Kaggle,
				},
				Skype: Student.SocialProfile{
					URL: oldStudent.Skype,
				},
			},

			StructVersion: 2,
			CreatedAt:     primitive.NewDateTimeFromTime(time.Now()),
			UpdatedAt:     primitive.NewDateTimeFromTime(time.Now()),
			DataErrors:    errorArray,
		}

		filter := bson.D{{Key: "_id", Value: oldStudent.ID}}

		if _, errUpdate := studentCollection.ReplaceOne(ctx, filter, newStudent); errUpdate != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"count": count, "error": errUpdate.Error(), "id": oldStudent.ID})
			return
		}

		count = count + 1
	}

	ctx.JSON(200, gin.H{"count": count})
}
