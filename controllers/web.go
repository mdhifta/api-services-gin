package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Web(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{"title": "Welcome to my structur"})
}
