package main

import (
	"fmt"

	"github.com/wahbifadhillah/learn-go/tree/main/06-structs-and-custom-types/user"
)

type str string

func (text str) log() {
	fmt.Println(text)
}

func main() {

	var log str = "this is custom log"
	log.log()
	UserFirstName := getUserData("Please enter your first name: ")
	UserLastName := getUserData("Please enter your last name: ")
	UserBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	// appUser = &user.User{

	// }
	appUser, err := user.New(UserFirstName, UserLastName, UserBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@t.t", "test123")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	// ... do something awesome with that gathered data!

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
