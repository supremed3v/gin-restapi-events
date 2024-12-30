package routes

import (
	"example.com/event-app/models"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(401, err)
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(404, err)
		return
	}

	ctx.JSON(201, gin.H{
		"message": "User created successfully",
	})
}
