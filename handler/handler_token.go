package handler

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"io/ioutil"

	"github.com/Eric-GreenComb/x-server/ether"
	"github.com/Eric-GreenComb/x-server/token"
)

// DeployToken DeployToken
func DeployToken(c *gin.Context) {

	_name := c.PostForm("name")
	_symbol := c.PostForm("symbol")
	_keystore := c.PostForm("keystore")
	_pwd := c.PostForm("pwd")

	keyjson, err := ioutil.ReadFile(_keystore)
	if err != nil {
		c.JSON(200, "keystore file is not exit")
		return
	}

	txOpt, err := bind.NewTransactor(strings.NewReader(string(keyjson)), _pwd)
	if err != nil {
		c.String(200, err.Error())
		return
	}

	_initialAmount := big.NewInt(1000000000000)
	_address, _, _, err := token.DeployHumanStandardToken(txOpt, ether.GetEthClient(), _initialAmount, _name, 10, _symbol)
	if err != nil {
		c.String(200, err.Error())
		return
	}

	fmt.Println(_address)

	address := fmt.Sprintf("0x%x", _address)

	c.JSON(200, address)
}

// BalanceOfToken BalanceOfToken
func BalanceOfToken(c *gin.Context) {

	_conaddr, _ := c.GetQuery("conaddr")
	_addr, _ := c.GetQuery("addr")

	_caller, err := token.NewHumanStandardTokenCaller(common.HexToAddress(_conaddr), ether.GetEthClient())
	if err != nil {
		c.String(200, err.Error())
		return
	}
	_bigint, err := _caller.BalanceOf(&bind.CallOpts{Pending: true}, common.HexToAddress(_addr))
	if err != nil {
		c.String(200, err.Error())
		return
	}
	c.JSON(200, _bigint.String())
}

// TransferToken TransferToken
func TransferToken(c *gin.Context) {

	_conaddr := c.PostForm("conaddr")
	_keystore := c.PostForm("keystore")
	_to := c.PostForm("to")
	_amount := c.PostForm("amount")
	_pwd := c.PostForm("pwd")

	keyjson, err := ioutil.ReadFile(_keystore)
	if err != nil {
		c.JSON(200, "keystore file is not exit")
		return
	}

	_int64, err := strconv.ParseInt(_amount, 10, 64)
	if err != nil {
		c.String(200, err.Error())
		return
	}

	txOpt, err := bind.NewTransactor(strings.NewReader(string(keyjson)), _pwd)
	if err != nil {
		c.String(200, err.Error())
		return
	}
	ts, _ := token.NewHumanStandardTokenTransactor(common.HexToAddress(_conaddr), ether.GetEthClient())

	_bigint := big.NewInt(_int64)
	_, err = ts.Transfer(txOpt, common.HexToAddress(_to), _bigint)
	if err != nil {
		c.String(200, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"from":   "from",
		"to":     _to,
		"amount": _amount,
	})
}
