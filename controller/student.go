package controller

import (
	"sort"
	"strings"

	"github.com/FrosTiK-SD/auth/constants"
	"github.com/FrosTiK-SD/auth/model"
	"github.com/FrosTiK-SD/auth/util"
	"github.com/FrosTiK-SD/models/misc"
	studentModel "github.com/FrosTiK-SD/models/student"
	db "github.com/FrosTiK-SD/mongik/db"
	models "github.com/FrosTiK-SD/mongik/models"
	"github.com/google/go-cmp/cmp"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func getAliasEmailList(email string) []string {
	var aliasEmailList []string
	aliasEmailList = append(aliasEmailList, email)
	aliasEmailList = append(aliasEmailList, strings.ReplaceAll(email, "iitbhu.ac.in", "itbhu.ac.in"))
	aliasEmailList = append(aliasEmailList, strings.ReplaceAll(email, "itbhu.ac.in", "iitbhu.ac.in"))
	sort.Strings(aliasEmailList)
	return aliasEmailList
}

func GetUserByEmail(mongikClient *models.Mongik, email *string, role *string, noCache bool) (*model.StudentPopulated, *string) {
	var studentPopulated model.StudentPopulated

	// Gets the alias emails
	emailList := getAliasEmailList(*email)

	// Query to DB
	studentPopulated, _ = db.AggregateOne[model.StudentPopulated](mongikClient, constants.DB, constants.COLLECTION_STUDENT, []bson.M{{
		"$match": bson.M{"email": bson.M{"$in": emailList}},
	}, {
		"$lookup": bson.M{
			"from":         constants.COLLECTION_GROUP,
			"localField":   "groups",
			"foreignField": "_id",
			"as":           "groups",
		},
	}}, noCache)

	// Now check if it is actually a student by the ROLES
	if !util.CheckRoleExists(&studentPopulated.GroupDetails, *role) {
		return nil, &constants.ERROR_NOT_A_STUDENT
	}

	return &studentPopulated, nil
}

func AssignUnVerifiedFields(updated *model.StudentPopulated, current *studentModel.Student) {
	// cannot change: groups, companies, batch, email, department, academics.verification, socialProfile.verification, metadata
	current.RollNo = updated.RollNo
	current.Course = updated.Course
	current.Specialisation = updated.Specialisation
	current.FirstName = updated.FirstName
	current.MiddleName = updated.MiddleName
	current.LastName = updated.LastName
	current.ProfilePicture = updated.ProfilePicture
	current.Gender = updated.Gender
	current.DOB = updated.DOB
	current.PermanentAddress = updated.PermanentAddress
	current.PresentAddress = updated.PresentAddress
	current.PersonalEmail = updated.PersonalEmail
	current.Mobile = updated.Mobile
	current.Category = updated.Category
	current.MotherTongue = updated.MotherTongue
	current.ParentsDetails = updated.ParentsDetails
}

func InvalidateVerifiedFieldsOnChange(updated *model.StudentPopulated, current *studentModel.Student) {
	// invalidate academic details
	if !cmp.Equal(updated.Academics, current.Academics) {
		current.Academics = updated.Academics
		current.Academics.Verification = misc.Verification{}
	}

	// invalidate social profiles
	if !cmp.Equal(updated.SocialProfiles.LinkedIn, current.SocialProfiles.LinkedIn) {
		current.SocialProfiles.LinkedIn = updated.SocialProfiles.LinkedIn
		current.SocialProfiles.LinkedIn.Verification = misc.Verification{}
	}

	if !cmp.Equal(updated.SocialProfiles.Github, current.SocialProfiles.Github) {
		current.SocialProfiles.Github = updated.SocialProfiles.Github
		current.SocialProfiles.Github.Verification = misc.Verification{}
	}

	if !cmp.Equal(updated.SocialProfiles.MicrosoftTeams, current.SocialProfiles.MicrosoftTeams) {
		current.SocialProfiles.MicrosoftTeams = updated.SocialProfiles.MicrosoftTeams
		current.SocialProfiles.MicrosoftTeams.Verification = misc.Verification{}
	}

	if !cmp.Equal(updated.SocialProfiles.Skype, current.SocialProfiles.Skype) {
		current.SocialProfiles.Skype = updated.SocialProfiles.Skype
		current.SocialProfiles.Skype.Verification = misc.Verification{}
	}

	if !cmp.Equal(updated.SocialProfiles.GoogleScholar, current.SocialProfiles.GoogleScholar) {
		current.SocialProfiles.GoogleScholar = updated.SocialProfiles.GoogleScholar
		current.SocialProfiles.GoogleScholar.Verification = misc.Verification{}
	}

	if !cmp.Equal(updated.SocialProfiles.Codeforces, current.SocialProfiles.Codeforces) {
		current.SocialProfiles.Codeforces = updated.SocialProfiles.Codeforces
		current.SocialProfiles.Codeforces.Verification = misc.Verification{}
	}

	if !cmp.Equal(updated.SocialProfiles.CodeChef, current.SocialProfiles.CodeChef) {
		current.SocialProfiles.CodeChef = updated.SocialProfiles.CodeChef
		current.SocialProfiles.CodeChef.Verification = misc.Verification{}
	}

	if !cmp.Equal(updated.SocialProfiles.LeetCode, current.SocialProfiles.LeetCode) {
		current.SocialProfiles.LeetCode = updated.SocialProfiles.LeetCode
		current.SocialProfiles.LeetCode.Verification = misc.Verification{}
	}

	if !cmp.Equal(updated.SocialProfiles.Kaggle, current.SocialProfiles.Kaggle) {
		current.SocialProfiles.Kaggle = updated.SocialProfiles.Kaggle
		current.SocialProfiles.Kaggle.Verification = misc.Verification{}
	}

	var newWorkExperienceArray []studentModel.WorkExperience

	// invalidate work experience
	for _, updatedWorkExp := range updated.WorkExperience {
		isUpdated := true
		for _, currentWorkExp := range current.WorkExperience {
			if cmp.Equal(updatedWorkExp, currentWorkExp) {
				newWorkExperienceArray = append(newWorkExperienceArray, currentWorkExp)
				isUpdated = false
				break
			}
		}

		if isUpdated {
			updatedWorkExp.Verification = misc.Verification{}
			newWorkExperienceArray = append(newWorkExperienceArray, updatedWorkExp)
		}
	}

	current.WorkExperience = newWorkExperienceArray

	// invalidate extra
	if !cmp.Equal(updated.Extras, current.Extras) {
		updated.Extras = current.Extras
		updated.Extras.Verification = misc.Verification{}
	}
}
