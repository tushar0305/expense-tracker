package models

import (
	"time"
	"fmt"
	"github.com/tushar0305/expense-tracker/db"
)

type Expense struct{
	Id			int64 		`json:"id"`
	Amount		int64		`json:"amount" binding:"required"`
	Category	string		`json:"category" binding:"required"`
	Description	string		`json:"description" binding:"required"`
	Date		time.Time	`json:"date" binding:"required"`
	UserId		int64		`json:"userId"`
}

func (e *Expense) Save() error {
    if db.Db == nil {
        return fmt.Errorf("database connection is not initialized")
    }

    query := `INSERT INTO expenses (amount, category, date, description, userId)
    VALUES(?, ?, ?, ?, ?)
    `

    stmt, err := db.Db.Prepare(query)
    if err != nil {
        fmt.Println("Error preparing query:", err)
        return err
    }
    defer stmt.Close()

    result, err := stmt.Exec(e.Amount, e.Category, e.Date, e.Description, e.UserId)
    if err != nil {
        fmt.Println("Error executing query:", err)
        return err
    }
    id, err := result.LastInsertId()
    if err != nil {
        fmt.Println("Error getting last insert id:", err)
        return err
    }
    e.Id = id

    return nil
}

