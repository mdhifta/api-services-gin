package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func MiddlewareApi(context *gin.Context) {
	jwtToken := context.GetHeader("Authorization")

	// validate
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECERET")), nil
	})

	// check error
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":    404,
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})

		context.Abort()

		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// set email from jwt
		context.Set("email", claims["Email"])

		// continue
		context.Next()
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":    404,
			"status":  false,
			"message": "Upps.. invalid to claim token.",
			"data":    nil,
		})

		context.Abort()

		return
	}

}
