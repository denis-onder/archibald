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
	return User{
		FirstName:   firstName,
		LastName:    lastName,
		DisplayName: buildDisplayName(firstName, lastName),
		Email:       email,
		Age:         age,
	}
}

func TestUserStruct(t *testing.T) {
	user := buildUser()

	if user.FirstName == "" {
		t.Errorf("First name is required")
	}

	if user.LastName == "" {
		t.Errorf("Last name is required")
	}

	if user.DisplayName == "" {
		t.Errorf("Display name is required")
	}

	if user.Email == "" {
		t.Errorf("Email is required")
	}

	if !strings.Contains(user.Email, "@mail.com") {
		t.Errorf("Email should be valid")
	}

	if user.Age < 0 {
		t.Errorf("Age must be a valid integer")
	}

	fmt.Print(user)
}
