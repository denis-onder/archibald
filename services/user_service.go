package services

import (
	. "github.com/denis-onder/archibald/builders"
	. "github.com/denis-onder/archibald/db"
	. "github.com/denis-onder/archibald/domain"
	. "github.com/denis-onder/archibald/validators"
)

type UserServiceType struct {
	CreateUser func(firstName, lastName, email string, age int) (error, User)
	GetUser    func(id int) (error, User)
	UpdateUser func(id int, updateUserObject User) bool
	DeleteUser func(id int) bool
}

func CreateUser(firstName, lastName, email string, age int) (error, User) {
	user := UserBuilder(firstName, lastName, email, age)
	isValid := ValidateUser(user)

	// TODO: Insert the user in the database if valid

	if isValid == nil {
		statement := `INSERT INTO users ("firstName", "lastName", "email", "age") VALUES ($1, $2, $3, $4)`
	}

	return isValid, user
}

func GetUser(id int) (error, User) {
	// Database query
	// If no user found return error
	// Return user
}

func UpdateUser(id int, updateUserObject User) bool {
	// Database query
	// If no user found return error
	// Validate user with modified data
	// If validation error, return error
	// Return user
}

func DeleteUser(id int) bool {
	// Database query
	// If no user found return false
	// Return true
}

var UserService UserServiceType = UserServiceType{
	CreateUser: CreateUser,
	GetUser:    GetUser,
	UpdateUser: UpdateUser,
	DeleteUser: DeleteUser,
}
