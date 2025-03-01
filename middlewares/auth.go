package middlewares

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/tushar0305/expense-tracker/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized!"})
		context.Abort()
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized!"})
		context.Abort()
		return
	}

	context.Set("userId", userId)
	context.Next()
}
