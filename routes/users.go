package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	_"github.com/tushar0305/expense-tracker/db"
	"github.com/tushar0305/expense-tracker/models"
	"github.com/tushar0305/expense-tracker/utils"
)

func SignUp(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	_, err = user.Save()
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not save User!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "user created sucessfully"})

}

func Login(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	err = user.ValidateCred()
	if err != nil{
		if err.Error() == "user not found" {
			context.JSON(http.StatusNotFound, gin.H{"message": "User not found!"})
			return
		} else if err.Error() == "invalid credentials" {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password!"})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong!"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"could not authenticate user!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login Successful!",
		"token": token,
		"user_id": user.Id,
	})
}