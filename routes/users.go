package routes

import (
	"courseProject/models"
	"courseProject/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse the data"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"messge": "Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"mesage": "User created successfully"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse the data1"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Couldn't authenticate user.2"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't authenticate user.3"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
