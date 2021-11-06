package domain

import (
	"strings"
	"testing"
)

var id int = 0
var firstName string = "Test"
var lastName string = "User"
var email string = "test@mail.com"
var password string = "1234567890"

func buildUser() User {
	user := User{}

	user.SetId(id)
	user.SetFirstName(firstName)
	user.SetLastName(lastName)
	user.SetDisplayName(firstName, lastName)
	user.SetEmail(email)
	user.SetPassword(password)

	return user
}

func TestUserStruct(t *testing.T) {
	user := buildUser()

	if user.GetId() < 0 {
		t.Errorf("An ID is required")
	}

	if user.GetFirstName() == "" {
		t.Errorf("First name is required")
	}

	if user.GetLastName() == "" {
		t.Errorf("Last name is required")
	}

	if user.GetDisplayName() == "" {
		t.Errorf("Display name is required")
	}

	if user.GetEmail() == "" {
		t.Errorf("Email is required")
	}

	if !strings.Contains(user.GetEmail(), "@mail.com") {
		t.Errorf("Email should be valid")
	}

	if user.GetPassword() == "" {
		t.Errorf("Password is required")
	}
}
