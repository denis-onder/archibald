package builders

import (
	"testing"
)

func TestUserBuilder(t *testing.T) {
	firstName := "Test"
	lastName := "User"
	email := "test@user.com"
	age := 20
	user := UserBuilder(firstName, lastName, email, age)

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

	if user.GetAge() != age {
		t.Errorf("Invalid age")
	}
}
