package controllers

import (
	"api-services/initializers"
	"api-services/models"
	"api-services/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Register(context *gin.Context) {
	// Get data from post
	var request struct {
		Fname        string
		Lname        string
		Phone_number string
		Password     string
		Email        string
	}

	context.Bind(&request)

	password, _ := utils.HashPassword(request.Password)

	// Create a post
	create := models.Users{
		Fname: request.Fname,
		Lname: request.Lname,

		Phone_number: request.Phone_number,
		Email:        request.Email,
		Password:     password,
		Roles:        1,
	}

	result := initializers.DB.Create(&create)

	// Return data
	if result.Error != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":    404,
			"status":  false,
			"message": "Upps.. error to create users register!",
			"data":    create,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"status":  true,
			"message": "Success.. create users register!",
			"data":    create,
		})
	}

	return
}

func Login(context *gin.Context) {
	// Get data from post
	var request struct {
		Email    string
		Password string
	}

	context.Bind(&request)

	// get data users to roles 1 (users)
	var users models.Users
	res := initializers.DB.First(&users, "email = ? AND roles = ?", request.Email, 1)

	if res.Error != nil {
		// create jwt
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    users.User_id,                                // create user id
			"email": users.Email,                                  // create user email
			"exp":   time.Now().Add(time.Hour * 24 * 9999).Unix(), // not ex token
		})

		// get string token
		tokenString, err := token.SignedString([]byte(os.Getenv("SECERET")))

		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code":    404,
				"status":  false,
				"message": "Upps.. invalid token account.",
				"data":    nil,
			})

			return
		}

		// check password is true or not
		if utils.CheckPasswordHash(request.Password, users.Password) {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"status":  true,
				"message": "Success login to apps.",
				"data":    users,
				"token":   tokenString,
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"code":    404,
				"status":  false,
				"message": "Upps.. your password wrong!",
				"data":    nil,
			})
		}

		return
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"status":  false,
			"message": "Upps.. your email not found!",
			"data":    nil,
		})
	}

	return
}
