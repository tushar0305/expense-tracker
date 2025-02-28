package models

import (
	"database/sql"
	"errors"
	"log"
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
	if err != nil{
		return 0, err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil{
		return 0, err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil{
		log.Fatalf("User Signup Failed: %v", err)
	}

	userId, err := result.LastInsertId()
	if err != nil{
		return 0, nil
	}

	return userId, nil
	
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

    passwordIsValid := utils.CheckHashedPassword(u.Password, retrievedPassword)
    if !passwordIsValid {
        return errors.New("invalid credentials")
    }

    return nil
}