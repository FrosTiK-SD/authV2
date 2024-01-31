package handler

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/FrosTiK-SD/auth/model"
	Student "github.com/FrosTiK-SD/auth/model"
	Constant "github.com/FrosTiK-SD/models/constant"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		Category: "",
		IsPWD:    false,
		IsEWS:    false,
	}

	if input == "" {
		return category
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

func GetXXIIYear(cursor *mongo.Cursor) (int, int) {
	var type1 model.XXIIYearType1
	var type2 model.XXIIYearType2

	if err1 := cursor.Decode(&type1); err1 == nil {
		xyear, errx := strconv.Atoi(type1.XYear)
		if errx != nil {
			xyear = 0
		}
		xiiyear, errxii := strconv.Atoi(type1.XiiYear)
		if errxii != nil {
			xiiyear = 0
		}

		return xyear, xiiyear
	}

	if err1 := cursor.Decode(&type2); err1 == nil {
		return type2.XYear, type2.XiiYear
	}

	return -1, -1
}

func (h *Handler) MigrateStudentDataToV2(ctx *gin.Context) {
	studentCollection := h.MongikClient.MongoClient.Database(Constant.DB).Collection(Constant.StudentCollection)
	cursor, errFind := studentCollection.Find(ctx, bson.D{{Key: "version", Value: bson.D{{Key: "$exists", Value: false}}}})
	if errFind != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"count": 0, "error": errFind.Error(), "reason": "Find not successful"})
		return
	}

	var count int64 = 0

	for cursor.Next(ctx) {
		var oldStudent model.OldStudent
		if errDecode := cursor.Decode(&oldStudent); errDecode != nil {
			fmt.Println(errDecode)
			continue
		}

		if oldStudent.Email == "tpo@itbhu.ac.in" {
			continue
		}

		if oldStudent.CompaniesAlloted == nil {
			oldStudent.CompaniesAlloted = []string{}
		}

		errorArray := []string{}
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
		case "btech":
			EndYearOffset = 4
			Course = Constant.BTECH
		default:
			EndYearOffset = -1
			Course = Constant.BTECH
			errorArray = append(errorArray, "course")
		}

		category := GetCategoryFromString(oldStudent.Category)

		kaggle := ""
		if oldStudent.Kaggle != "" {
			kaggle = oldStudent.Kaggle
		}
		if oldStudent.Kaggel != "" {
			kaggle = oldStudent.Kaggel
		}

		jeeRank, errJeeRank := GetRankFromString(oldStudent.JeeRank, category)
		if errJeeRank != nil {
			fmt.Println(errJeeRank)
			errorArray = append(errorArray, "jeeRank")
			jeeRank.Rank = -1
		}

		educationGap, err := strconv.Atoi(regexp.MustCompile(`\d+`).FindString(oldStudent.EducationGap))
		if err != nil {
			fmt.Println(educationGap)
			errorArray = append(errorArray, "educationGap")
			educationGap = -1
		}

		xYear, xiiYear := GetXXIIYear(cursor)
		if xYear <= 0 {
			errorArray = append(errorArray, "xYear")
		}

		if xiiYear <= 0 {
			errorArray = append(errorArray, "xiiYear")
		}

		dob, errDOB := GetDOBFromString(oldStudent.Dob)
		if errDOB != nil {
			fmt.Println(errDOB)
			errorArray = append(errorArray, "dob")
			dob = 0
		}

		gender := Constant.Gender(strings.ToLower(oldStudent.Gender))

		newStudent := Student.Student{
			Id:               oldStudent.ID,
			Groups:           oldStudent.Groups,
			CompaniesAlloted: oldStudent.CompaniesAlloted,

			Batch: &Student.Batch{
				StartYear: oldStudent.Batch,
				EndYear:   oldStudent.Batch + EndYearOffset,
			},
			RollNo:         oldStudent.RollNo,
			InstituteEmail: oldStudent.Email,
			Department:     oldStudent.Department,
			Course:         &Course,

			FirstName: oldStudent.FirstName,
			LastName:  oldStudent.LastName,

			Gender:           &gender,
			DOB:              &dob,
			PermanentAddress: oldStudent.PermanentAddress,
			PresentAddress:   oldStudent.PresentAddress,
			PersonalEmail:    oldStudent.PersonalEmail,
			Mobile:           strconv.FormatInt(oldStudent.Mobile, 10),
			Category:         &category,
			MotherTongue:     oldStudent.MotherTongue,
			ParentsDetails: &Student.ParentsDetails{
				FatherName:       oldStudent.FatherName,
				FatherOccupation: oldStudent.FatherOccupation,
				MotherName:       oldStudent.MotherName,
				MotherOccupation: oldStudent.MotherOccupation,
			},

			Academics: Student.Academics{
				JEERank: &jeeRank,
				XthClass: &Student.EducationDetails{
					Certification: oldStudent.XBoard,
					Institute:     oldStudent.XInstitute,
					Year:          xYear,
					Score:         oldStudent.XPercentage,
				},
				XIIthClass: &Student.EducationDetails{
					Certification: oldStudent.XiiBoard,
					Institute:     oldStudent.XiiInstitute,
					Year:          xiiYear,
					Score:         oldStudent.XiiPercentage,
				},
				EducationGap: educationGap,
				SemesterSPI: Student.SemesterSPI{
					One:   oldStudent.SemesterOne,
					Two:   oldStudent.SemesterTwo,
					Three: oldStudent.SemesterThree,
					Four:  oldStudent.SemesterFour,
					Five:  oldStudent.SemesterFive,
					Six:   oldStudent.SemesterSix,
					Seven: oldStudent.SemesterSeven,
					Eight: oldStudent.SemesterEight,
				},
				SummerTermSPI: Student.SummerTermSPI{
					One:   oldStudent.SummerOne,
					Two:   oldStudent.SummerTwo,
					Three: oldStudent.SummerThree,
					Four:  oldStudent.SummerFour,
					Five:  oldStudent.SummerFive,
				},
				UnderGraduate: &Student.EducationDetails{
					Certification: oldStudent.UgIn,
					Institute:     oldStudent.UgCollege,
					Year:          oldStudent.UgYear,
					Score:         oldStudent.UgScore,
				},
				PostGraduate: &Student.EducationDetails{
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
				LinkedIn: &Student.SocialProfile{
					URL: oldStudent.LinkedIn,
				},
				MicrosoftTeams: &Student.SocialProfile{
					URL: oldStudent.MicrosoftTeams,
				},
				Github: &Student.SocialProfile{
					URL: oldStudent.Github,
				},
				Kaggle: &Student.SocialProfile{
					URL: kaggle,
				},
				Skype: &Student.SocialProfile{
					URL: oldStudent.Skype,
				},
			},
			Extras: Student.Extras{
				VideoResume: &oldStudent.VideoResume,
			},

			StructVersion: 2,
			CreatedAt:     primitive.NewDateTimeFromTime(time.Now()),
			UpdatedAt:     primitive.NewDateTimeFromTime(time.Now()),
			DataErrors:    errorArray,
		}

		ValidateData(&newStudent)

		filter := bson.D{{Key: "_id", Value: oldStudent.ID}}

		if _, errUpdate := studentCollection.ReplaceOne(ctx, filter, newStudent); errUpdate != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"count": count, "error": errUpdate.Error(), "id": oldStudent.ID})
			return
		}

		count = count + 1
	}

	ctx.JSON(200, gin.H{"count": count})
}
