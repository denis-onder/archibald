package controllers

import (
	. "github.com/denis-onder/archibald/domain"
	UserService "github.com/denis-onder/archibald/services"
)

func CreateUser(firstName, lastName, email string, age int) User {
	err, user := UserService.CreateUser(firstName, lastName, email, age)

	if err != nil {
		// Handle error
	}

	return user
}

func GetUser(id int) User {
	err, user := UserService.GetUser(id)

	if err != nil {
		// Handle error
	}

	return user
}

func UpdateUser(user User) bool {
	//
}

func DeleteUser(id int) bool {
	return UserService.DeleteUser(id)
}
