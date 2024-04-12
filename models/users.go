package models

import (
	"courseProject/db"
	"courseProject/utils"
	"errors"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO users(email, password)
	VALUES
	(?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.Id = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := `
	SELECT id, password FROM users WHERE email = ?
	`
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.Id, &retrievedPassword)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
