package services

import (
	. "github.com/denis-onder/archibald/builders"
	. "github.com/denis-onder/archibald/domain"
	. "github.com/denis-onder/archibald/validators"
)

type UserServiceType struct {
	CreateUser func(firstName, lastName, email string, password string) (User, error)
	GetUser    func(id int) (User, error)
	UpdateUser func(id int, updateUserObject User) bool
	DeleteUser func(id int) bool
}

func CreateUser(firstName, lastName, email string, password string) (User, error) {
	user := UserBuilder(firstName, lastName, email, password)
	isValid := ValidateUser(user)

	// TODO: Insert the user in the database if valid

	if isValid != nil {
		// Handle error
	}

	statement := `INSERT INTO users ("firstName", "lastName", "email", "password") VALUES ($1, $2, $3, $4)`

	res, err := Database.Exec(statement, firstName, lastName, email, password)

	if err != nil {
		// handle error
	}

	// extract the user from the result

	return res, isValid
}

func GetUser(id int) (User, error) {
	statement := `SELECT * FROM users WHERE id='$1'`

	// Database query
	res, err := Database.Exec(statement, id)

	if err != nil {
		// handle error
	}

	// Return user
	return res, err
}

func UpdateUser(id int, updateUserObject User) bool {
	// Database query
	// If no user found return error
	// Validate user with modified data
	// If validation error, return error
	return false
}

func DeleteUser(id int) bool {
	statement := `DELETE FROM users WHERE id='$1'`
	_, err := Database.Exec(statement, id)

	if err != nil {
		return false
	}

	return true
}

var UserService UserServiceType = UserServiceType{
	CreateUser: CreateUser,
	GetUser:    GetUser,
	UpdateUser: UpdateUser,
	DeleteUser: DeleteUser,
}
