package routes

import (
	"log"
	"net/http"
	"strconv"

	"example.com/event-app/models"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Couldn't fetch event",
			"error":   err,
		})
		return
	}
	ctx.JSON(200, events)
}

func getSingleEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Couldn't fetch event",
			"error":   err,
		})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Couldn't fetch event",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, event)

}

func createEvent(ctx *gin.Context) {

	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(401, err)
		return
	}

	userId := ctx.GetInt64("userId")

	event.UserID = userId

	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Couldn't create event",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "event created",
		"event":   event,
	})

}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Couldn't parse event id",
			"error":   err,
		})
		return
	}

	_, err = models.GetEventByID(eventId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Couldn't fetch event",
			"error":   err.Error(),
		})
		return
	}

	var updatedEvent models.Event

	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		log.Printf("Error binding JSON: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Couldn't parse event JSON",
			"error":   err.Error(),
		})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Message": "Couldn't update event during process",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "event updated succesfully",
	})

}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Couldn't parse event id",
			"error":   err,
		})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Couldn't fetch event",
			"error":   err.Error(),
		})
		return
	}

	err = event.Delete()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Couldn't delete event",
			"error":   err.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"message": "event deleted successfully",
	})
}
