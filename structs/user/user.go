package user

import (
	"fmt"
	"time"
	"errors"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
	CreatedAt time.Time
}

func (u User) OutputUserDetails() {
	fmt.Println(u.FirstName, u.LastName, u.Birthdate)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}

func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birth date are required")
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}, nil
}
