package controllers

import (
	. "github.com/denis-onder/archibald/domain"
	UserService "github.com/denis-onder/archibald/services"
)

func CreateUser(firstName, lastName, email string, password string) (User, error) {
	user, err := UserService.CreateUser(firstName, lastName, email, password)

	if err != nil {
		// Handle error
		return user, err
	}

	return user, nil
}

func GetUser(id int) User {
	user, err := UserService.GetUser(id)

	if err != nil {
		// Handle error
	}

	return user
}

func UpdateUser(user User) bool {
	return false
}

func DeleteUser(id int) bool {
	return UserService.DeleteUser(id)
}
