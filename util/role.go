package util

import "frostik.com/auth/model"

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
