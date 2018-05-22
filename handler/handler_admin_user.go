package handler

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/x-server/bean"
	"github.com/Eric-GreenComb/x-server/persist"
	"github.com/Eric-GreenComb/x-server/regexp"
)

// CreateAdminUser CreateAdminUser
func CreateAdminUser(c *gin.Context) {

	_userID := c.PostForm("userID")
	_name := c.PostForm("name")
	_passwd := c.PostForm("passwd")
	_email := c.PostForm("email")

	if _userID == "" || _name == "" || _passwd == "" {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "There are some empty fields."})
		return
	}

	if !regexp.IsMobile(_userID) {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "UserID must phone number."})
		return
	}

	sum := sha256.Sum256([]byte(_passwd))
	_Passwd := fmt.Sprintf("%x", sum)

	var _user bean.AdminUsers
	_user.UserID = _userID
	_user.Name = _name
	_user.Passwd = _Passwd
	_user.Email = _email

	err := persist.GetPersist().CreateAdminUser(_user)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _user})
	}
}

// AdminUserInfo AdminUserInfo
func AdminUserInfo(c *gin.Context) {

	_userid := c.Params.ByName("userid")

	user, err := persist.GetPersist().AdminUserInfo(_userid)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": user})
}

// UpdateAdminUserPasswd UpdateAdminUserPasswd
func UpdateAdminUserPasswd(c *gin.Context) {

	_userid := c.Params.ByName("userid")
	_old := c.Params.ByName("old")
	_new := c.Params.ByName("new")

	sumOld := sha256.Sum256([]byte(_old))
	oldPasswd := fmt.Sprintf("%x", sumOld)

	_, err := persist.GetPersist().AdminLogin(_userid, oldPasswd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	sum := sha256.Sum256([]byte(_new))
	newPasswd := fmt.Sprintf("%x", sum)

	err = persist.GetPersist().UpdateAdminUserPasswd(_userid, newPasswd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": "success"})
}
