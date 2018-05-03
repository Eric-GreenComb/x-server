package handler

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/Eric-GreenComb/x-server/bean"
	"github.com/Eric-GreenComb/x-server/persist"
	"github.com/Eric-GreenComb/x-server/regexp"
	"github.com/gin-gonic/gin"
)

// CreateUser Create User
func CreateUser(c *gin.Context) {

	user := bean.Users{}
	c.Bind(&user)

	if user.UserID == "" || user.Passwd == "" || user.Name == "" {
		c.JSON(422, gin.H{"errcode": 1, "msg": "There are some empty fields."})
		return
	}

	if !regexp.IsMobile(user.UserID) {
		c.JSON(422, gin.H{"errcode": 1, "msg": "UserID must phone number."})
		return
	}

	sum := sha256.Sum256([]byte(user.Passwd))
	user.Passwd = fmt.Sprintf("%x", sum)

	err := persist.GetPersist().CreateUser(user)

	if err != nil {
		c.JSON(422, gin.H{"errcode": 1, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": user})
	}
}

// UserInfo User Info
func UserInfo(c *gin.Context) {

	_userid := c.Params.ByName("userid")

	user, err := persist.GetPersist().UserInfo(_userid)

	if err != nil {
		c.JSON(422, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": user})
}
