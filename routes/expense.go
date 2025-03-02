package routes

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/tushar0305/expense-tracker/models"
)

func CreateExpense(context *gin.Context) {
	var expense models.Expense

	err := context.ShouldBindJSON(&expense)
if err != nil {
    fmt.Println("JSON Binding Error:", err)
    context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
    return
}


	userId, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized"})
		return
	}
	expense.UserId = userId.(int64)

	err = expense.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save expense"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message":    "Expense created successfully",
		"expense_id": expense.Id,
	})
}

func GetExpenses(context *gin.Context) {
	userId, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized"})
		return
	}

	userID := userId.(int64)

	// Default range last 30 days
	startDate := time.Now().AddDate(0, 0, -30)
	endDate := time.Now()

	// Checking for filters in query parameters
	dateRange := context.Query("range")
	startParam := context.Query("start")
	endParam := context.Query("end")

	switch dateRange {
	case "week":
		startDate = time.Now().AddDate(0, 0, -7)
	case "month":
		startDate = time.Now().AddDate(0, -1, 0)
	case "3months":
		startDate = time.Now().AddDate(0, -3, 0)
	}

	// If custom range is provided, override defaults
	if startParam != "" && endParam != "" {
		parsedStart, err1 := time.Parse("2006-01-02", startParam)
		parsedEnd, err2 := time.Parse("2006-01-02", endParam)
		if err1 == nil && err2 == nil {
			startDate = parsedStart
			endDate = parsedEnd
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid date format. Use YYYY-MM-DD."})
			return
		}
	}

	// Fetching expenses from database
	expenses, err := models.GetExpensesByUser(userID, startDate, endDate)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch expenses", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, expenses)
}

func UpdateExpense(context *gin.Context) {

}

func DeleteExpense(context *gin.Context) {

}

