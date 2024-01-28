package util

import "github.com/FrosTiK-SD/authV2/model"

func CheckRoleExists(groups *[]model.Group, role string) bool {
	checkRoleStatus := false
	for _, group := range *groups {
		if ArrayContains(group.Roles, role) {
			checkRoleStatus = true
			break
		}
	}
	return checkRoleStatus
}
