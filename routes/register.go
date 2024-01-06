package routes

import (
	"net/http"
	"strconv"

	"github.com/criticalsession/go-rest/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	eventId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Id", "error": err.Error()})
		return
	}

	ev, err := models.GetEventById(uint(eventId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event", "error": err.Error()})
		return
	}

	err = ev.Register(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user to event", "error": err.Error()})
		return
	}

	ev.Registrations = append(ev.Registrations, models.Registration{EventId: uint(eventId), UserId: userId})

	ctx.JSON(http.StatusCreated,
		gin.H{
			"message": "User registered to event",
			"event":   ev,
			"userId":  userId,
		})

}

func unregisterFromEvent(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	eventId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Id", "error": err.Error()})
		return
	}

	ev, err := models.GetEventById(uint(eventId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event", "error": err.Error()})
		return
	}

	err = ev.Unregister(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not unregister user to event", "error": err.Error()})
		return
	}

	registrations := []models.Registration{}
	for _, r := range ev.Registrations {
		if r.UserId != userId {
			registrations = append(registrations, r)
		}
	}
	ev.Registrations = registrations

	ctx.JSON(http.StatusCreated,
		gin.H{
			"message": "User unregistered from event",
			"event":   ev,
			"userId":  userId,
		})
}
