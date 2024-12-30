package routes

import (
	"example.com/event-app/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getSingleEvent)
	autheticated := server.Group("/")
	autheticated.Use(middlewares.Authenticate)
	autheticated.POST("/events", createEvent)
	autheticated.PUT("/events/:id", updateEvent)
	autheticated.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
