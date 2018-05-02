package handler

import (
	"net/http"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// GetHello GetHello
func GetHello(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(http.StatusOK, gin.H{
		"userID": claims["id"],
		"text":   "Get Hello",
	})
}

// PostHello PostHello
func PostHello(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(http.StatusOK, gin.H{
		"userID": claims["id"],
		"text":   "Post Hello",
	})
}
