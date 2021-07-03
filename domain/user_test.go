package domain

import (
	"fmt"
	"strings"
	"testing"
)

var firstName string = "Test"
var lastName string = "User"
var displayName string = buildDisplayName(firstName, lastName)
var email string = "test@mail.com"
var age = 21

func buildDisplayName(firstName, lastName string) string {
	builder := strings.Builder{}

	builder.WriteString(firstName)
	builder.WriteString(" ")
	builder.WriteString(lastName)

	return builder.String()
}

func buildUser() User {
	user := User{}

	user.setFirstName(firstName)
	user.setLastName(lastName)
	user.setDisplayName(displayName)
	user.setEmail(email)
	user.setAge(age)

	return user
}

func TestUserStruct(t *testing.T) {
	user := buildUser()

	if user.getFirstName() == "" {
		t.Errorf("First name is required")
	}

	if user.getLastName() == "" {
		t.Errorf("Last name is required")
	}

	if user.getDisplayName() == "" {
		t.Errorf("Display name is required")
	}

	if user.getEmail() == "" {
		t.Errorf("Email is required")
	}

	if !strings.Contains(user.getEmail(), "@mail.com") {
		t.Errorf("Email should be valid")
	}

	if user.getAge() < 0 {
		t.Errorf("Age must be a valid integer")
	}

	fmt.Print(user)
}
