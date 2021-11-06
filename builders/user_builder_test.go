package builders

import (
	"testing"
)

func TestUserBuilder(t *testing.T) {
	firstName := "Test"
	lastName := "User"
	email := "test@user.com"
	password := "1234567890"
	user := UserBuilder(firstName, lastName, email, password)

	if user.GetFirstName() != firstName {
		t.Errorf("Invalid first name")
	}

	if user.GetLastName() != lastName {
		t.Errorf("Invalid last name")
	}

	if user.GetDisplayName() != firstName+" "+lastName {
		t.Errorf("Invalid display name")
	}

	if user.GetEmail() != email {
		t.Errorf("Invalid email address")
	}

	if user.GetPassword() == "" {
		t.Errorf("A password is required")
	}

	if user.GetPassword() == password {
		t.Errorf("The password is not encrypted")
	}
}
