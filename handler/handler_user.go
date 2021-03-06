package handler

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/x-server/bean"
	"github.com/Eric-GreenComb/x-server/persist"
	"github.com/Eric-GreenComb/x-server/regexp"
)

// CreateUser Create User
func CreateUser(c *gin.Context) {

	_userID := c.PostForm("userID")
	_name := c.PostForm("name")
	_passwd := c.PostForm("passwd")
	_email := c.PostForm("email")
	fmt.Println("userID : " + _userID)

	if _userID == "" || _name == "" || _passwd == "" {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "There are some empty fields."})
		return
	}

	if !regexp.IsMobile(_userID) {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "UserID must phone number."})
		return
	}

	fmt.Println("create user")

	sum := sha256.Sum256([]byte(_passwd))
	_Passwd := fmt.Sprintf("%x", sum)

	var _user bean.Users
	_user.UserID = _userID
	_user.Name = _name
	_user.Passwd = _Passwd
	_user.Email = _email

	err := persist.GetPersist().CreateUser(_user)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _user})
	}
}

// UserInfo User Info
func UserInfo(c *gin.Context) {

	_userid := c.Params.ByName("userid")

	user, err := persist.GetPersist().UserInfo(_userid)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": user})
}

// UpdateUserPasswd Update User Passwd
func UpdateUserPasswd(c *gin.Context) {

	_userid := c.Params.ByName("userid")
	_old := c.Params.ByName("old")
	_new := c.Params.ByName("new")

	sumOld := sha256.Sum256([]byte(_old))
	oldPasswd := fmt.Sprintf("%x", sumOld)

	_, err := persist.GetPersist().Login(_userid, oldPasswd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	sum := sha256.Sum256([]byte(_new))
	newPasswd := fmt.Sprintf("%x", sum)

	err = persist.GetPersist().UpdateUserPasswd(_userid, newPasswd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": "success"})
}

// CountUser User Count
func CountUser(c *gin.Context) {

	_count, err := persist.GetPersist().CountUser()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _count})
}

// ListUser ListUser
func ListUser(c *gin.Context) {

	_search := c.Params.ByName("search")
	_page := c.Params.ByName("page")
	_nPage, _ := strconv.Atoi(_page)

	_users, err := persist.GetPersist().ListUser(_search, _nPage)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "get users error"})
		return
	}

	c.JSON(http.StatusOK, _users)
}
