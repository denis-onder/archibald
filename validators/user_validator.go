package validators

import (
	"strings"

	. "github.com/denis-onder/archibald/domain"
)

func ValidateUser(u User) bool {
	if u.GetFirstName() == "" {
		return false
	}

	if u.GetLastName() == "" {
		return false
	}

	if u.GetDisplayName() == "" {
		return false
	}

	if u.GetEmail() == "" {
		return false
	}

	if !strings.Contains(u.GetEmail(), "@mail.com") {
		return false
	}

	if u.GetAge() < 0 {
		return false
	}

	return true
}
