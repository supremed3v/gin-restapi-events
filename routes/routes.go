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
	autheticated.POST("/events/:id/register", registerForEvent)
	autheticated.DELETE("/events/:id/register", cancelRegistration)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
