package routes

import (
	"fmt"
	"net/http"
	"strconv"
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

func UpdateExpenseById(context *gin.Context) {
	userId, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized"})
		return
	}
	userID := userId.(int64)

	expenseID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid expense ID"})
		return
	}

	existingExpense, err := models.GetExpenseByID(expenseID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Expense not found"})
		return
	}

	if existingExpense.UserId != userID {
		context.JSON(http.StatusForbidden, gin.H{"message": "Not authorized to update this expense"})
		return
	}

	var updatedExpense models.Expense
	if err := context.ShouldBindJSON(&updatedExpense); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body", "error": err.Error()})
		return
	}

	existingExpense.Amount = updatedExpense.Amount
	existingExpense.Category = updatedExpense.Category
	existingExpense.Description = updatedExpense.Description
	existingExpense.Date = updatedExpense.Date

	if err := existingExpense.Update(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update expense"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Expense updated successfully"})

}

func DeleteExpenseById(context *gin.Context) {
	userId, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized"})
		return
	}
	userID := userId.(int64)

	expenseID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid expense ID"})
		return
	}

	existingExpense, err := models.GetExpenseByID(expenseID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Expense not found"})
		return
	}
	if existingExpense.UserId != userID {
		context.JSON(http.StatusForbidden, gin.H{"message": "Not authorized to delete this expense"})
		return
	}

	if err := models.DeleteExpense(expenseID); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete expense"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Expense deleted successfully"})
}

