package routes

import (
	"net/http"

	"github.com/criticalsession/go-rest/models"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
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
