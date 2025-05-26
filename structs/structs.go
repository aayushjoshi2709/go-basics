package main

import (
	"fmt"
	"time"
	"example.com/structs/user"
)


func main() {
	userFirstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	// null struct
	// appUser = user {}

	appUser := user.User{
		FirstName: userFirstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}

	// this can be done if you assign in same order
	// defined in blueprint
	// appUser = user{
	// 	userFirstName,
	// 	lastName,
	// 	birthdate,
	// 	time.Now(),
	// }

	// ... do something awesome with that gathered data!

	outputUserDetails(appUser)
	outputUserDetailsPtr(&appUser)
	appUser.OutputUserDetails()

	appUser.ClearUserName()
	appUser.OutputUserDetails()

	appUser2, err := user.NewUser("Mamta", "Joshi", "09/06/2002")
	if err != nil {
		fmt.Println(err)
		return
	}
	appUser2.OutputUserDetails()

	appUser3, err := user.NewUser("", "", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	appUser3.OutputUserDetails()

}

func outputUserDetails(u user.User) {
	fmt.Println(u.FirstName, u.LastName, u.Birthdate)
}

func outputUserDetailsPtr(u *user.User) {
	// u.lastName = (*u).firstName for structs in go
	fmt.Println((*u).FirstName, u.LastName, u.Birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
