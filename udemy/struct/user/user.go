package user

import (
	"errors"
	"fmt"
	"time"
)

// struct
type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// struct embedding
type Admin struct {
	email    string
	password string
	User
}

// just a pattern in go to create constructor
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("all fields are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
		},
	}
}

// changing the normal function into method
// (user User) is called receiver argument
func (user User) OutputUserData() {
	fmt.Println(user.firstName)
}

func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}
