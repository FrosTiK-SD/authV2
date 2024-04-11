package util

import (
	"strings"

	"github.com/FrosTiK-SD/auth/constants"
)

func CheckValidInstituteEmail(email string) bool {
	_, domain, found := strings.Cut(email, "@")
	if found && ArrayContains(constants.INSTITUTE_MAIL_DOMAINS, domain) {
		return true
	}

	return false
}
