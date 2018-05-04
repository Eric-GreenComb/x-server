package handler

import (
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/x-server/bean"
	"github.com/Eric-GreenComb/x-server/persist"
)

const (
	veryLightScryptN = 2
	veryLightScryptP = 1
)

// CreateAccount CreateAccount
func CreateAccount(c *gin.Context) {

	_userID := c.Params.ByName("userid")
	_password := c.Params.ByName("password")

	_, ks := tmpKeyStore(true)
	a, err := ks.NewAccount(_password)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": "NewAccount error"})
		return
	}

	keyjson, err := ioutil.ReadFile(a.URL.Path)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": "file is not exit"})
		return
	}
	var _address bean.Addresses
	_address.UserID = _userID
	_address.Address = a.Address.String()
	_address.KeyStore = string(keyjson)

	err = persist.GetPersist().CreateAddress(_address)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	//删除keystore文件
	err = os.Remove(a.URL.Path)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(200, a.Address)
}

// GetKeystore GetKeystore
func GetKeystore(c *gin.Context) {

	_address := c.Params.ByName("address")

	_keystore, err := persist.GetPersist().AddressInfo(_address)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": "create address error"})
		return
	}

	c.String(200, _keystore.KeyStore)
}

func tmpKeyStore(encrypted bool) (string, *keystore.KeyStore) {
	d := "./tmp"

	new := keystore.NewPlaintextKeyStore
	if encrypted {
		new = func(kd string) *keystore.KeyStore {
			return keystore.NewKeyStore(kd, veryLightScryptN, veryLightScryptP)
		}
	}
	return d, new(d)
}
