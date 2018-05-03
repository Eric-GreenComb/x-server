package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHello GetHello
func GetHello(c *gin.Context) {
	_userID := c.MustGet("userID")
	c.JSON(http.StatusOK, gin.H{
		"sub":  _userID,
		"text": "Get Hello",
	})
}

// PostHello PostHello
func PostHello(c *gin.Context) {
	_userID := c.MustGet("userID")
	c.JSON(http.StatusOK, gin.H{
		"sub":  _userID,
		"text": "Post Hello",
	})
}
