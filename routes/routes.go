package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tushar0305/expense-tracker/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", SignUp)
    server.POST("/login", Login)

	authenticated := server.Group("/")
    authenticated.Use(middlewares.Authenticate)

	authenticated.GET("/expenses", GetExpenses)
	authenticated.POST("/expense", CreateExpense)
	authenticated.PUT("/expense/:id", UpdateExpenseById)
	authenticated.DELETE("/expense/:id", DeleteExpenseById)
}