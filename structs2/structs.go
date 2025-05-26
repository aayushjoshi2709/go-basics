package main

import (
	"fmt"

	"example.com/structs2/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser3, err := user.New(userFirstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}
	appUser3.OutputUserDetails()

	admin := user.NewAdmin("test@gmail.com", "Tinatu@123")
	admin.MyUser.OutputUserDetails()

	// this works because of
	// type Admin struct {
	// email    string
	// password string
	// MyUser   User
	// User <---
	admin.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
