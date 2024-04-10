package util

import (
	"strings"

	"github.com/FrosTiK-SD/auth/constants"
)

func CheckValidInstituteEmail(email string) bool {
	_, domain, found := strings.Cut(email, "@")
	if found && domain == constants.INSTITUTE_MAIL_DOMAIN {
		return true
	}

	return false
}
