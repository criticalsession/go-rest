package routes

import (
	"net/http"

	"github.com/criticalsession/go-rest/models"
	"github.com/criticalsession/go-rest/utils"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated,
		gin.H{
			"message": "User created",
			"user":    user,
		})
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed", "error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated,
		gin.H{
			"message": "Login successful",
			"token":   token,
		})
}
