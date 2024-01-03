package main

import (
	"net/http"

	"github.com/criticalsession/go-rest/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.GetAllEvents())
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
	}

	event.Id = 1
	event.UserId = 1
	event.Save()

	ctx.JSON(http.StatusCreated,
		gin.H{
			"message": "Event created",
			"event":   event,
		})
}
