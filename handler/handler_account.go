package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/x-server/bean"
	"github.com/Eric-GreenComb/x-server/ether"
	"github.com/Eric-GreenComb/x-server/persist"
)

// CreateAccount CreateAccount
func CreateAccount(c *gin.Context) {

	_userID := c.Params.ByName("userid")
	_password := c.Params.ByName("password")

	_key, err := ether.Ks.NewKey()
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	keyjson, err := ether.Ks.GenKeystore(_key, _password)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}
	var _address bean.Addresses
	_address.UserID = _userID
	_address.Address = _key.Address.String()
	_address.KeyStore = string(keyjson)

	err = persist.GetPersist().CreateAddress(_address)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(200, _key.Address.String())
}

// ListAccount ListAccount
func ListAccount(c *gin.Context) {

	_userID := c.Params.ByName("userid")

	_addresses, err := persist.GetPersist().ListAddress(_userID)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": "get address error"})
		return
	}

	c.JSON(200, _addresses)
}

// GetKeystore GetKeystore
func GetKeystore(c *gin.Context) {

	_address := c.Params.ByName("address")

	_keystore, err := persist.GetPersist().AddressInfo(_address)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": "get address error"})
		return
	}

	c.String(200, _keystore.KeyStore)
}

// UpdateAccountPwd UpdateAccountPwd
func UpdateAccountPwd(c *gin.Context) {

	_addr := c.Params.ByName("addr")
	_password := c.Params.ByName("password")
	_newpassword := c.Params.ByName("newpassword")

	_keystore, err := persist.GetPersist().AddressInfo(_addr)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": "get address error"})
		return
	}

	keyjson, err := ether.Ks.Update([]byte(_keystore.KeyStore), _password, _newpassword)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	err = persist.GetPersist().UpdateAccountPwd(_keystore.UserID, _addr, string(keyjson))
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(200, _addr)
}
