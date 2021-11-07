package controllers

import (
	"strings"
	"testing"
)

var userId int

func TestCreateUser(t *testing.T) {
	firstName := "Test"
	lastName := "User"
	email := "test@user.com"
	password := "1234567890"

	user, err := CreateUser(firstName, lastName, email, password)

	if err != nil {
		t.Errorf("An error has occured")
	}

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

func TestGetUser(t *testing.T) {
}

func TestUpdateUser(t *testing.T) {
}

func TestDeleteUser(t *testing.T) {
}
