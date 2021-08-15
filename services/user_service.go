package services

import (
	. "github.com/denis-onder/archibald/builders"
	. "github.com/denis-onder/archibald/domain"
	. "github.com/denis-onder/archibald/validators"
)

func CreateUser(firstName, lastName, email string, age int) (error, User) {
	user := UserBuilder(firstName, firstName, email, age)
	isValid := ValidateUser(user)

	// TODO: Insert the user in the database if valid

	return isValid, user
}

func GetUser(id int) User {
}

func UpdateUser(id int, updateUserObject User) bool {
}

func DeleteUser(id int) bool {
}
