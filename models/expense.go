package models

import (
	"fmt"
	"time"
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

func GetExpensesByUser(userID int64, startDate time.Time, endDate time.Time) ([]Expense, error) {
	if db.Db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query := `SELECT id, amount, category, date, description, userId 
			  FROM expenses 
			  WHERE userId = ? AND date BETWEEN ? AND ? 
			  ORDER BY date DESC`

	rows, err := db.Db.Query(query, userID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var expense Expense
		err := rows.Scan(&expense.Id, &expense.Amount, &expense.Category, &expense.Date, &expense.Description, &expense.UserId)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}


