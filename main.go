package main

import (
	"net/http"

	"example.com/event-app/db"
	"example.com/event-app/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080

}

func getEvents(ctx *gin.Context) {
	events := models.GetAllEvents()
	ctx.JSON(200, events)
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

	event.Save()

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "event created",
		"event":   event,
	})

}
