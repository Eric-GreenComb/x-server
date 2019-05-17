package handler

import (
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/x-server/bean"
	"github.com/Eric-GreenComb/x-server/ether"
	"github.com/Eric-GreenComb/x-server/persist"
	"github.com/Eric-GreenComb/x-server/token"
)

// CreateUserAddressTokens CreateUserAddressTokens User
func CreateUserAddressTokens(c *gin.Context) {

	_userID := c.Params.ByName("userid")
	_address := c.Params.ByName("address")
	_tokenaddress := c.Params.ByName("tokenaddress")

	_, err := persist.GetPersist().GetUserAddressTokens(_userID, _address, _tokenaddress)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": "OK"})
		return
	}

	token, err := persist.GetPersist().TokenInfo(_tokenaddress)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	var userAddressToken bean.UserAddressTokens
	userAddressToken.UserID = _userID
	userAddressToken.Address = _address
	userAddressToken.ESN = token.ESN
	userAddressToken.TokenAddress = _tokenaddress
	userAddressToken.Name = token.Name
	userAddressToken.Symbol = token.Symbol

	err = persist.GetPersist().CreateUserAddressTokens(userAddressToken)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": userAddressToken})
	}
}

// DeleteUserAddressTokens DeleteUserAddressTokens Info
func DeleteUserAddressTokens(c *gin.Context) {

	_userID := c.Params.ByName("userid")
	_address := c.Params.ByName("address")
	_tokenaddress := c.Params.ByName("tokenaddress")

	err := persist.GetPersist().DeleteUserAddressTokens(_userID, _address, _tokenaddress)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": "success"})
}

// GetUserAddressTokens GetUserAddressTokens Info
func GetUserAddressTokens(c *gin.Context) {

	_userID := c.Params.ByName("userid")
	_address := c.Params.ByName("address")
	_tokenaddress := c.Params.ByName("tokenaddress")

	userAddressTokens, err := persist.GetPersist().GetUserAddressTokens(_userID, _address, _tokenaddress)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": userAddressTokens})
}

// ListUserAddressTokens ListUserAddressTokens Info
func ListUserAddressTokens(c *gin.Context) {

	_userID := c.Params.ByName("userid")
	_address := c.Params.ByName("address")

	userAddressTokens, err := persist.GetPersist().ListUserAddressTokens(_userID, _address)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": userAddressTokens})
}

// ListUserAddressTokenBanlance ListUserAddressTokenBanlance Info
func ListUserAddressTokenBanlance(c *gin.Context) {

	_userID := c.Params.ByName("userid")
	_address := c.Params.ByName("address")

	userAddressTokens, err := persist.GetPersist().ListUserAddressTokens(_userID, _address)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	var tokenValues []bean.TokenValue
	for _, userAddressToken := range userAddressTokens {

		_caller, err := token.NewHumanStandardTokenCaller(common.HexToAddress(userAddressToken.TokenAddress), ether.GetEthClient())
		if err != nil {
			fmt.Println("Caller Error : " + err.Error())
			continue
		}
		_bigint, err := _caller.BalanceOf(&bind.CallOpts{Pending: false}, common.HexToAddress(_address))
		if err != nil {
			fmt.Println("BalanceOf Error : " + err.Error())
			continue
		}

		var tokenValue bean.TokenValue
		tokenValue.Address = userAddressToken.TokenAddress
		tokenValue.Name = userAddressToken.Name
		tokenValue.Symbol = userAddressToken.Symbol
		tokenValue.Balance = _bigint.String()

		tokenValues = append(tokenValues, tokenValue)
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": tokenValues})
}
