package builders

import . "github.com/denis-onder/archibald/domain"

func UserBuilder(firstName, lastName, email string, age int) User {
	user := User{}

	user.SetFirstName(firstName)
	user.SetLastName(lastName)
	user.SetDisplayName(firstName, lastName)
	user.SetEmail(email)
	user.SetAge(age)

	return user
}
