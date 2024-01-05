package routes

import (
	"net/http"
	"strconv"

	"github.com/criticalsession/go-rest/models"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	res, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func getEvent(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Id", "error": err.Error()})
		return
	}

	event, err := models.GetEventById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
	}

	event.UserId = 1
	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated,
		gin.H{
			"message": "Event created",
			"event":   event,
		})
}

func updateEvent(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Id", "error": err.Error()})
		return
	}

	_, err = models.GetEventById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event", "error": err.Error()})
		return
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
	}
	updatedEvent.Id = uint(id)

	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event", "error": err.Error()})
	}

	ctx.JSON(http.StatusCreated,
		gin.H{
			"message": "Event updated",
			"event":   updatedEvent,
		})
}

func deleteEvent(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Id", "error": err.Error()})
		return
	}

	e, err := models.GetEventById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event", "error": err.Error()})
		return
	}

	err = e.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event", "error": err.Error()})
	}

	ctx.JSON(http.StatusCreated,
		gin.H{
			"message": "Event deleted",
			"eventId": id,
		})
}
