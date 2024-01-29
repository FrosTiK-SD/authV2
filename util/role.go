package util

import (
	student "github.com/FrosTiK-SD/models/company"
)

func CheckRoleExists(groups *[]student.Group, role string) bool {
	checkRoleStatus := false
	for _, group := range *groups {
		if ArrayContains(group.Roles, role) {
			checkRoleStatus = true
			break
		}
	}
	return checkRoleStatus
}
