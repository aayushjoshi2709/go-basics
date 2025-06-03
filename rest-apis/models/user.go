package models

import (
	"aayushjoshi2709/rest-apis/db"
	"aayushjoshi2709/rest-apis/utils"
	"errors"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := `INSERT INTO users(email, password) values (?, ?);`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, user.Password)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	user.ID = userId
	return err
}

func (user *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, user.Email)
	var retrivedPassword string
	err := row.Scan(&user.ID, &retrivedPassword)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrivedPassword)

	if !passwordIsValid {
		return errors.New("invalid password")
	}
	return nil
}
