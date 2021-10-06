package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AccessMiddleware struct {
	Roles    []string
	Accessor string
}

func (am *AccessMiddleware) Check(c *gin.Context) {
	authHeader := strings.Split(c.Request.Header.Get("Authorization"), " ")

	if len(authHeader) > 2 {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"data":  "",
			"error": "Wrong token format",
		})
		return
	}
	if authHeader[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"data":  "",
			"error": "Wrong token format (Should start with 'Bearer')",
		})
		return
	}

	am.Accessor = authHeader[1]
	claims := jwtgo.MapClaims{}
	_, err := jwtgo.ParseWithClaims(am.Accessor, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"data":  "",
			"error": err.Error(),
		})
		return
	}

	for _, value := range am.Roles {
		if claims["role"] == value {
			c.Next()
			return
		}
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"data":  "",
		"error": "You must be an admin to access this route",
	})
	return
}
