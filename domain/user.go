package domain

import "strings"

type User struct {
	id          int
	firstName   string
	lastName    string
	displayName string
	email       string
	password    string
}

func buildDisplayName(firstName, lastName string) string {
	builder := strings.Builder{}

	builder.WriteString(firstName)
	builder.WriteString(" ")
	builder.WriteString(lastName)

	return builder.String()
}

// Getters
func (u User) GetId() int {
	return u.id
}

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

func (u User) GetPassword() string {
	return u.password
}

// Setters
func (u *User) SetId(id int) {
	u.id = id
}

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

func (u *User) SetPassword(password string) {
	u.password = password
}
