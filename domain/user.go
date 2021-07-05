package domain

import "strings"

type User struct {
	firstName   string
	lastName    string
	displayName string
	email       string
	age         int
}

func buildDisplayName(firstName, lastName string) string {
	builder := strings.Builder{}

	builder.WriteString(firstName)
	builder.WriteString(" ")
	builder.WriteString(lastName)

	return builder.String()
}

// Getters
func (u User) GetFirstName() string {
	return u.firstName
}

func (u User) GetLastName() string {
	return u.lastName
}

func (u User) GetDisplayName() string {
	return u.displayName
}

func (u User) GetEmail() string {
	return u.email
}

func (u User) GetAge() int {
	return u.age
}

// Setters
func (u *User) SetFirstName(firstName string) {
	u.firstName = firstName
}

func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}

func (u *User) SetDisplayName(firstName, lastName string) {
	u.displayName = buildDisplayName(firstName, lastName)
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) SetAge(age int) {
	u.age = age
}
