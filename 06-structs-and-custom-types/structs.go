package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func (u *user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func newUser(firstName, lastName, birthDate string) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, last name and birthdate are required")
	}
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	UserFirstName := getUserData("Please enter your first name: ")
	UserLastName := getUserData("Please enter your last name: ")
	UserBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user
	appUser, err := newUser(UserFirstName, UserLastName, UserBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ... do something awesome with that gathered data!

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
