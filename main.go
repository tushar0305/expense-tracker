package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tushar0305/expense-tracker/db"
	"github.com/tushar0305/expense-tracker/routes"
)

func main(){
	db.InitDB()
	log.Println("Database initialized successfully")

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}