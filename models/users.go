package models

import (
	"errors"
	"database/sql"
	"github.com/tushar0305/expense-tracker/db"
	"github.com/tushar0305/expense-tracker/utils"
)

type User struct{
	Id 			int64	`json:"id"`
	Email 		string	`json:"email" binding:"required"`
	Password 	string	`json:"password" binding:"required"`
}

func (u *User) Save() (int64, error) {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`

	stmt, err := db.Db.Prepare(query)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil{
		return 0, err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return 0, err
	}

	UserId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	u.Id = UserId
	return UserId, nil
}

func (u *User) ValidateCred() error {
	query := `SELECT id, email, password FROM users WHERE email = ?`
	row := db.Db.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.Id, &u.Email, &retrievedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user not found")
		}
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}