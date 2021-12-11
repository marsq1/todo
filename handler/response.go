package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Error struct {
	Message string
}

func NewErrorResponse (ctx *gin.Context, statusCode int, message string) {
	log.Println(message)
	ctx.AbortWithStatusJSON(statusCode, Error{Message: message})
}
