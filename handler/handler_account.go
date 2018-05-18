package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/x-server/bean"
	"github.com/Eric-GreenComb/x-server/config"
	"github.com/Eric-GreenComb/x-server/ether"
	"github.com/Eric-GreenComb/x-server/persist"
)

// CreateAccount CreateAccount
func CreateAccount(c *gin.Context) {

	_userID := c.Params.ByName("userid")
	_name := c.Params.ByName("name")
	_password := c.Params.ByName("password")

	_key, err := ether.Ks.NewKey()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	keyjson, err := ether.Ks.GenKeystore(_key, _password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}
	backjson, err := ether.Ks.GenKeystore(_key, config.ServerConfig.Passphrase)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}
	var _address bean.Addresses
	_address.UserID = _userID
	_address.Name = _name
	_address.Address = _key.Address.String()
	_address.KeyStore = string(keyjson)
	_address.BackStore = string(backjson)

	err = persist.GetPersist().CreateAddress(_address)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, _key.Address.String())
}

// ListAccount ListAccount
func ListAccount(c *gin.Context) {

	_userID := c.Params.ByName("userid")

	_addresses, err := persist.GetPersist().ListAddress(_userID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "get address error"})
		return
	}

	c.JSON(http.StatusOK, _addresses)
}

// GetKeystore GetKeystore
func GetKeystore(c *gin.Context) {

	_address := c.Params.ByName("address")

	_keystore, err := persist.GetPersist().AddressInfo(_address)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "get address error"})
		return
	}

	c.String(http.StatusOK, _keystore.KeyStore)
}

// UpdateAccountPwd UpdateAccountPwd
func UpdateAccountPwd(c *gin.Context) {

	_addr := c.Params.ByName("addr")
	_password := c.Params.ByName("password")
	_newpassword := c.Params.ByName("newpassword")

	_keystore, err := persist.GetPersist().AddressInfo(_addr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "get address error"})
		return
	}

	keyjson, err := ether.Ks.Update([]byte(_keystore.KeyStore), _password, _newpassword)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	err = persist.GetPersist().UpdateAccountPwd(_keystore.UserID, _addr, string(keyjson))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, _addr)
}

// RecoverAccountPwd RecoverAccountPwd
func RecoverAccountPwd(c *gin.Context) {

	_addr := c.Params.ByName("addr")
	_newpassword := c.Params.ByName("newpassword")

	_keystore, err := persist.GetPersist().AddressInfo(_addr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "get address error"})
		return
	}

	keyjson, err := ether.Ks.Update([]byte(_keystore.BackStore), config.ServerConfig.Passphrase, _newpassword)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	err = persist.GetPersist().UpdateAccountPwd(_keystore.UserID, _addr, string(keyjson))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, _addr)
}
