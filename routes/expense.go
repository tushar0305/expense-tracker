package routes

import (
	"fmt"
	"net/http"
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

func ListExpense(context *gin.Context) {

}

func UpdateExpense(context *gin.Context) {

}

func DeleteExpense(context *gin.Context) {

}

