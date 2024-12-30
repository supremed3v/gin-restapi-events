package routes

import (
	"example.com/event-app/models"
	"example.com/event-app/utils"
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

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(402, err)
		return
	}

	err = user.Login()

	if err != nil {
		ctx.JSON(401, err.Error())
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		ctx.JSON(401, err.Error())
		return
	}

	ctx.JSON(200, gin.H{"message": "login successfull", "token": token})

}
