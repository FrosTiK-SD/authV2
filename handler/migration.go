package handler

import (
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

func (h *Handler) MigrateStudentDataToV1Format(ctx *gin.Context) (int64, error) {
	studentCollection := h.MongikClient.MongoClient.Database(Constant.DB).Collection(Constant.StudentCollection)
	cursor, errFind := studentCollection.Find(ctx, bson.D{{Key: "version", Value: bson.D{{Key: "$exists", Value: false}}}})
	if errFind != nil {
		return 0, errFind
	}

	var count int64 = 0

	for cursor.Next(ctx) {
		var oldStudent model.OldStudent

		if errDecode := cursor.Decode(&oldStudent); errDecode != nil {
			return count, errDecode
		}

		var EndYearOffset int
		switch oldStudent.Course {
		case "idd":
			EndYearOffset = 5
		case "mtech":
			EndYearOffset = 2
		case "phd":
			EndYearOffset = -1
		default:
			EndYearOffset = 4
		}

		category := GetCategoryFromString(oldStudent.Category)

		jeeRank, errJeeRank := GetRankFromString(oldStudent.JeeRank, category)
		if errJeeRank != nil {
			return count, errJeeRank
		}
		xYear, errXYear := strconv.Atoi(oldStudent.XYear)
		if errXYear != nil {
			return count, errXYear
		}

		xiiYear, errXiiYear := strconv.Atoi(oldStudent.XiiYear)
		if errXiiYear != nil {
			return count, errXiiYear
		}

		dob, errDOB := GetDOBFromString(oldStudent.Dob)
		if errDOB != nil {
			return count, errDOB
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
			Course:         Constant.Course(oldStudent.Course),

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
					Score:         float32(oldStudent.XPercentage),
				},
				XIIthClass: Student.EducationDetails{
					Certification: oldStudent.XiiBoard,
					Institute:     oldStudent.XiiInstitute,
					Year:          xiiYear,
					Score:         float32(oldStudent.XiiPercentage),
				},
				EducationGap: -1,
				SemesterDetails: Student.SemesterSPI{
					One:   float32(oldStudent.SemesterOne),
					Two:   float32(oldStudent.SemesterTwo),
					Three: float32(oldStudent.SemesterThree),
					Four:  float32(oldStudent.SemesterFour),
					Five:  float32(oldStudent.SemesterFive),
					Six:   float32(oldStudent.SemesterSix),
				},
				SummerTermDetails: Student.SummerTermSPI{
					One:   float32(oldStudent.SummerOne),
					Two:   float32(oldStudent.SummerTwo),
					Three: float32(oldStudent.SummerThree),
					Four:  float32(oldStudent.SummerFour),
					Five:  float32(oldStudent.SummerFive),
				},
				CurrentCGPA:    float32(oldStudent.Cgpa),
				ActiveBacklogs: oldStudent.ActiveBacklogs,
				TotalBacklogs:  oldStudent.TotalBacklogs,
			},
			SocialProfiles: Student.SocialProfiles{
				LinkedIn: Student.SocialProfile{
					URL: oldStudent.LinkedIn,
				},
			},

			StructVersion: 2,
			CreatedAt:     primitive.NewDateTimeFromTime(time.Now()),
			UpdatedAt:     primitive.NewDateTimeFromTime(time.Now()),
		}

		filter := bson.D{{Key: "_id", Value: oldStudent.ID}}

		if _, errUpdate := studentCollection.ReplaceOne(ctx, filter, newStudent); errUpdate != nil {
			return count, errUpdate
		}

		count = count + 1
	}

	return count, nil
}
