package main

import (
	"net/http"
	"strconv"

	"example.com/event-app/db"
	"example.com/event-app/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getSingleEvent)
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080

}

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

	event.ID = 1
	event.UserID = 1

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
