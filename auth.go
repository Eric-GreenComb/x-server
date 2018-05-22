package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/x-server/bean"
	"github.com/Eric-GreenComb/x-server/persist"
)

// Login Login
func Login(c *gin.Context) {
	var _user bean.Users

	if c.Bind(&_user) != nil {
		AbortWithError(c, http.StatusBadRequest, "Missing usename or password", bean.Realm)
		return
	}

	fmt.Println(_user)

	sum := sha256.Sum256([]byte(_user.Passwd))
	_pwd := fmt.Sprintf("%x", sum)

	user, err := persist.GetPersist().Login(_user.UserID, _pwd)
	if err != nil {
		AbortWithError(c, http.StatusInternalServerError, "DB Query Error", bean.Realm)
		return
	}

	expire := time.Now().Add(bean.ExpireTime)
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	claims := make(jwt.MapClaims)
	claims["sub"] = user.UserID
	claims["exp"] = expire.Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(bean.JWTSigningKey))
	if err != nil {
		AbortWithError(c, http.StatusUnauthorized, "Create JWT Token faild", bean.Realm)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  tokenString,
		"expire": expire.Format(time.RFC3339),
	})
}

// AdminLogin AdminLogin
func AdminLogin(c *gin.Context) {
	var _adminUser bean.AdminUsers

	if c.Bind(&_adminUser) != nil {
		AbortWithError(c, http.StatusBadRequest, "Missing usename or password", bean.Realm)
		return
	}

	sum := sha256.Sum256([]byte(_adminUser.Passwd))
	_pwd := fmt.Sprintf("%x", sum)

	user, err := persist.GetPersist().AdminLogin(_adminUser.UserID, _pwd)
	if err != nil {
		AbortWithError(c, http.StatusInternalServerError, "DB Query Error", bean.Realm)
		return
	}

	expire := time.Now().Add(bean.ExpireTime)
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	claims := make(jwt.MapClaims)
	claims["sub"] = user.UserID
	claims["exp"] = expire.Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(bean.JWTSigningKey))
	if err != nil {
		AbortWithError(c, http.StatusUnauthorized, "Create JWT Token faild", bean.Realm)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  tokenString,
		"expire": expire.Format(time.RFC3339),
	})
}

// JWTAuth JWTAuth
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		_token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(bean.JWTSigningKey))
			return b, nil
		})

		if err != nil {
			AbortWithError(c, http.StatusUnauthorized, "Invaild User Token", bean.Realm)
			return
		}

		claims := _token.Claims.(jwt.MapClaims)

		c.Set("userID", claims["sub"])
	}
}

// RefreshToken RefreshToken
func RefreshToken(c *gin.Context) {

	_userID := c.MustGet("userID")
	expire := time.Now().Add(bean.ExpireTime)

	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["sub"] = _userID
	claims["exp"] = expire.Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	// Set some claims
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(bean.JWTSigningKey))

	if err != nil {
		AbortWithError(c, http.StatusUnauthorized, "Create JWT Token faild", bean.Realm)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  tokenString,
		"expire": expire.Format(time.RFC3339),
	})
}

// AbortWithError AbortWithError
func AbortWithError(c *gin.Context, code int, message, realm string) {
	c.Header("WWW-Authenticate", "JWT realm="+realm)
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
	c.Abort()
}
