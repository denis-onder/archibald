package domain

type User struct {
	firstName   string
	lastName    string
	displayName string
	email       string
	age         int
}

// Getters
func (u User) getFirstName() string {
	return u.firstName
}

func (u User) getLastName() string {
	return u.lastName
}

func (u User) getDisplayName() string {
	return u.displayName
}

func (u User) getEmail() string {
	return u.email
}

func (u User) getAge() int {
	return u.age
}

// Setters
func (u *User) setFirstName(firstName string) {
	u.firstName = firstName
}

func (u *User) setLastName(lastName string) {
	u.lastName = lastName
}

func (u *User) setDisplayName(displayName string) {
	u.displayName = displayName
}

func (u *User) setEmail(email string) {
	u.email = email
}

func (u *User) setAge(age int) {
	u.age = age
}
