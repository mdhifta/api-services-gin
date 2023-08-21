package controllers

import (
	"api-services/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* first you set data struct
model to get respon from api */
type dataList struct {
	Result   string `json:"result"`
	Response string `json:"response"`
	Data     struct {
		Id         string `json:"id"`
		Type       string `json:"tyoe"`
		Attributes struct {
			Name       string `json:"name"`
			Visibility string `json:"visibility"`
			Version    int    `json:"version"`
		}
		Relationships []struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		}
	} `json:"data"`
}

func GetResponse(context *gin.Context) {
	var responseApi = utils.RequestGET("list/e1d40cf9-33e9-4e80-a64c-7354d4c520d3")
	textResponse := []byte(responseApi)
	response := dataList{}

	err_response := json.Unmarshal(textResponse, &response)

	if err_response != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":    401,
			"status":  false,
			"message": err_response,
			"data":    responseApi,
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"status":  true,
		"message": "Success get data response",
		"data":    response,
	})

	return
}
