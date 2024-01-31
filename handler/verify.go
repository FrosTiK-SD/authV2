package handler

import (
	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/auth/controller"
	"github.com/FrosTiK-SD/auth/model"
	constant "github.com/FrosTiK-SD/models/constant"
	"github.com/FrosTiK-SD/models/misc"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HandlerVerifyStudentIdToken(ctx *gin.Context) {
	idToken := ctx.GetHeader("token")
	noCache := false
	if ctx.GetHeader("cache-control") == constants.NO_CACHE {
		noCache = true
	}

	email, exp, err := controller.VerifyToken(h.MongikClient.CacheClient, idToken, h.JwkSet, noCache)

	if err != nil {
		ctx.JSON(200, gin.H{
			"student": nil,
			"expire":  exp,
			"error":   err,
		})
	} else {
		student, err := controller.GetUserByEmail(h.MongikClient, email, &constants.ROLE_STUDENT, noCache)
		if h.Config.Mode == MIDDLEWARE {
			h.Session.Student = student
		} else {
			ctx.JSON(200, gin.H{
				"data":   student,
				"error":  err,
				"expire": exp,
			})
		}
	}

}

func (h *Handler) InvalidateCache(ctx *gin.Context) {
	h.MongikClient.CacheClient.Delete("GCP_JWKS")
	ctx.JSON(200, gin.H{
		"message": "Successfully invalidated cache",
	})
}

func ValidateEducationDetails(ed **model.EducationDetails) {
	if *ed != nil && (*ed).Score <= 0 {
		*ed = nil
	}
}

func ValidateRankDetails(rd **model.RankDetails) {
	if *rd != nil && (*rd).Rank <= 0 {
		*rd = nil
	}
}

func ValidateString(s **string) {
	if *s != nil && *(*s) == "" {
		*s = nil
	}
}

func ValidateSocialProfile(sp **model.SocialProfile) {
	if *sp != nil && (*sp).URL == "" {
		*sp = nil
	}
}

func ValidateAttachment(a **misc.Attachment) {
	if *a != nil && (*a).URL == "" {
		*a = nil
	}
}

func ValidateBatch(b **model.Batch) {
	if *b != nil && (*b).StartYear <= 0 {
		*b = nil
	}
}

func ValidateGender(g **constant.Gender) {
	if *g != nil && *(*g) == "" {
		*g = nil
	}
}

func ValidateCourse(c **constant.Course) {
	if *c != nil && *(*c) == "" {
		*c = nil
	}
}

func ValidateData(student *model.Student) {
	ValidateBatch(&student.Batch)
	ValidateCourse(&student.Course)
	ValidateString(&student.Specialisation)
	ValidateString(&student.MiddleName)
	ValidateAttachment(&student.ProfilePicture)
	ValidateGender(&student.Gender)

	if student.DOB != nil && *student.DOB <= 10 {
		student.DOB = nil
	}

	if student.Category.Category == "" {
		student.Category = nil
	}

	if student.ParentsDetails.FatherName == "" && student.ParentsDetails.FatherOccupation == "" && student.ParentsDetails.MotherName == "" && student.ParentsDetails.MotherOccupation == "" {
		student.ParentsDetails = nil
	}

	ValidateRankDetails(&student.Academics.JEERank)
	ValidateRankDetails(&student.Academics.GATERank)
	ValidateEducationDetails(&student.Academics.XthClass)
	ValidateEducationDetails(&student.Academics.XIIthClass)
	ValidateEducationDetails(&student.Academics.UnderGraduate)
	ValidateString(&student.Academics.Honours)
	ValidateEducationDetails(&student.Academics.PostGraduate)

	if student.Academics.ThesisEndDate != nil && *student.Academics.ThesisEndDate <= 0 {
		student.Academics.ThesisEndDate = nil
	}

	ValidateSocialProfile(&student.SocialProfiles.LinkedIn)
	ValidateSocialProfile(&student.SocialProfiles.Github)
	ValidateSocialProfile(&student.SocialProfiles.MicrosoftTeams)
	ValidateSocialProfile(&student.SocialProfiles.Skype)
	ValidateSocialProfile(&student.SocialProfiles.GoogleScholar)
	ValidateSocialProfile(&student.SocialProfiles.Codeforces)
	ValidateSocialProfile(&student.SocialProfiles.CodeChef)
	ValidateSocialProfile(&student.SocialProfiles.LeetCode)
	ValidateSocialProfile(&student.SocialProfiles.Kaggle)
}
