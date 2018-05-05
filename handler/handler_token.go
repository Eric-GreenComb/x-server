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

	"github.com/Eric-GreenComb/x-server/bean"
	"github.com/Eric-GreenComb/x-server/ether"
	"github.com/Eric-GreenComb/x-server/persist"
	"github.com/Eric-GreenComb/x-server/token"
)

// DeployToken DeployToken
func DeployToken(c *gin.Context) {

	_name := c.PostForm("name")
	_symbol := c.PostForm("symbol")
	_address := c.PostForm("address")
	_pwd := c.PostForm("pwd")

	_keystore, err := persist.GetPersist().AddressInfo(_address)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": "get address error"})
		return
	}

	txOpt, err := bind.NewTransactor(strings.NewReader(_keystore.KeyStore), _pwd)
	if err != nil {
		c.String(200, err.Error())
		return
	}

	_initialAmount := big.NewInt(1000000000000)
	_tokenAddress, _, _, err := token.DeployHumanStandardToken(txOpt, ether.GetEthClient(), _initialAmount, _name, 10, _symbol)
	if err != nil {
		c.String(200, err.Error())
		return
	}

	address := fmt.Sprintf("0x%x", _tokenAddress)

	c.JSON(200, address)
}

// BalanceOfToken BalanceOfToken
func BalanceOfToken(c *gin.Context) {

	_addr := c.Params.ByName("addr")
	_conaddrs := c.PostForm("conaddrs")

	split := strings.Split(_conaddrs, ",")

	var tokenValues []bean.TokenValue
	for _, _conaddr := range split {

		_caller, err := token.NewHumanStandardTokenCaller(common.HexToAddress(_conaddr), ether.GetEthClient())
		if err != nil {
			continue
		}
		_bigint, err := _caller.BalanceOf(&bind.CallOpts{Pending: true}, common.HexToAddress(_addr))
		if err != nil {
			continue
		}

		var tokenValue bean.TokenValue
		tokenValue.Address = _conaddr
		tokenValue.Balance = _bigint.String()

		tokenValues = append(tokenValues, tokenValue)
	}

	c.JSON(200, tokenValues)
}

// TransferToken TransferToken
func TransferToken(c *gin.Context) {

	_conaddr := c.PostForm("conaddr")
	_from := c.PostForm("from")
	_to := c.PostForm("to")
	_amount := c.PostForm("amount")
	_pwd := c.PostForm("pwd")

	_keystore, err := persist.GetPersist().AddressInfo(_from)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": "get address error"})
		return
	}

	_int64, err := strconv.ParseInt(_amount, 10, 64)
	if err != nil {
		c.String(200, err.Error())
		return
	}

	txOpt, err := bind.NewTransactor(strings.NewReader(_keystore.KeyStore), _pwd)
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

// CreateToken CreateToken
func CreateToken(c *gin.Context) {

	var _tokens bean.Tokens
	c.Bind(&_tokens)

	err := persist.GetPersist().CreateToken(_tokens)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": "create token error"})
		return
	}

	c.JSON(200, _tokens)
}

// TokenInfo TokenInfo
func TokenInfo(c *gin.Context) {

	_address := c.Params.ByName("address")

	token, err := persist.GetPersist().TokenInfo(_address)

	if err != nil {
		c.JSON(422, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": token})
}

// UpdateTokenWeight UpdateTokenWeight
func UpdateTokenWeight(c *gin.Context) {
	_address := c.Params.ByName("address")
	_weight := c.Params.ByName("weight")
	if _address == "" || _weight == "" {
		c.JSON(422, gin.H{"errcode": 1, "msg": "There are some empty fields."})
		return
	}

	_iWeight, err := strconv.Atoi("_weight")
	if err != nil {
		c.JSON(422, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	if err := persist.GetPersist().UpdateTokenWeight(_address, _iWeight); err != nil {
		c.JSON(406, gin.H{"errcode": 1, "msg": "update weight error."})
	} else {
		c.JSON(200, gin.H{"errcode": 0, "msg": "success"})
	}
}

// ListToken ListToken
func ListToken(c *gin.Context) {
	_page, _ := c.GetQuery("page") // page
	_nPage, _ := strconv.Atoi(_page)

	_tokens, err := persist.GetPersist().ListToken(_nPage)
	if err != nil {
		c.JSON(406, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _tokens})
}

// CreateTokenTransfer CreateTokenTransfer
func CreateTokenTransfer(c *gin.Context) {

	var _transfer bean.TokenTransfer
	c.Bind(&_transfer)

	err := persist.GetPersist().CreateTokenTransfer(_transfer)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 1, "msg": "create token error"})
		return
	}

	c.JSON(200, _transfer)
}

// ListTokenTransfer ListTokenTransfer
func ListTokenTransfer(c *gin.Context) {
	_address := c.Params.ByName("address")
	_page := c.Params.ByName("page")
	_nPage, _ := strconv.Atoi(_page)

	_transfers, err := persist.GetPersist().ListTokenTransfer(_address, _nPage)
	if err != nil {
		c.JSON(406, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _transfers})
}
