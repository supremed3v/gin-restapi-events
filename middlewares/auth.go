package middlewares

import (
	"net/http"

	"example.com/event-app/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Couldn't create event, please login",
		})
		return
	}
	uID, err := utils.VerifyToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(401, err)
		return
	}

	ctx.Set("userId", uID)

	ctx.Next()
}
