package validators

import (
	"errors"
	"strings"

	. "github.com/denis-onder/archibald/domain"
)

func ValidateUser(u User) error {
	if u.GetFirstName() == "" {
		return errors.New("First name is required.")
	}

	if u.GetLastName() == "" {
		return errors.New("Last name is required.")
	}

	if u.GetDisplayName() == "" {
		return errors.New("Display name is required.")
	}

	if u.GetEmail() == "" {
		return errors.New("An email address is required.")
	}

	if !strings.Contains(u.GetEmail(), "@mail.com") {
		return errors.New("Please provide a valid email address")
	}

	if u.GetAge() < 0 {
		return errors.New("Age must be a positive number")
	}

	return nil
}
