package main

import (
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/x-server/persist"
)

// AuthMiddleware the jwt middleware
var AuthMiddleware jwt.GinJWTMiddleware

// LoadAuthMiddleware load jwt
func LoadAuthMiddleware() jwt.GinJWTMiddleware {
	return jwt.GinJWTMiddleware{
		Realm:      "FiFu.io Blockchain",
		Key:        []byte("fifu.io blockchain"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {

			user, err := persist.GetPersist().Login(userId, password)
			if err != nil {
				return userId, false
			}

			return user.UserID, true
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			return true
			// if userId == "admin" {
			// 	return true
			// }

			// return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}
}