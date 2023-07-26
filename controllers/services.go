package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Services(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"status":  true,
		"message": "Yeay gin connected with api!",
		"data": map[string]interface{}{
			"access": true,
		},
	})
}
