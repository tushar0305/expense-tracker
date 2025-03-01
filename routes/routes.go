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

	authenticated.GET("/expenses", ListExpense)
	authenticated.POST("/expense", CreateExpense)
	authenticated.PUT("/expenses/:id", UpdateExpense)
	authenticated.DELETE("/expenses/:id", DeleteExpense)

}