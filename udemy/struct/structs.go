package main

import (
	"example.com/structs/user"
	"fmt"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// var appUser User

	//creating an instance of User
	// appUser := User{
	// 	firstName: firstName,
	// 	lastName:  lastName,
	// 	birthdate: birthdate,
	// 	createdAt: time.Now(),
	// }

	//calling constructor for creating instance of user(it is manually written constructor)
	appUser, err := user.New(firstName, lastName, birthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	adminUser := user.NewAdmin("test@ex.com", "password")
	fmt.Println(adminUser)

	// outputUserData(appUser) //normal function calling

	//calling method
	appUser.OutputUserData()
	appUser.ClearUserName()
	appUser.OutputUserData()

	// fmt.Println(firstName, lastName, birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

//this is normal function
// func outputUserData(user *User) {
// 	fmt.Println((*user).firstName)
// }
