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
	"github.com/Eric-GreenComb/x-server/config"
	"github.com/Eric-GreenComb/x-server/ether"
	"github.com/Eric-GreenComb/x-server/persist"
	"github.com/Eric-GreenComb/x-server/token"
)

// DeployToken DeployToken
func DeployToken(c *gin.Context) {

	_userID := c.PostForm("userID")

	_name := c.PostForm("name")
	_symbol := c.PostForm("symbol")
	_total := c.PostForm("total")
	_desc := c.PostForm("desc")

	_int64, err := strconv.ParseInt(_total, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_addresses, err := persist.GetPersist().ListAddress(_userID)
	if err != nil || len(_addresses) == 0 {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "get user address error"})
		return
	}

	fmt.Println(_addresses[0].Address)

	_keystore, err := persist.GetPersist().AddressInfo(_addresses[0].Address)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "get address error"})
		return
	}

	txOpt, err := bind.NewTransactor(strings.NewReader(_keystore.BackStore), config.ServerConfig.Passphrase)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_initialAmount := big.NewInt(_int64)
	_tokenAddress, _, _, err := token.DeployHumanStandardToken(txOpt, ether.GetEthClient(), _initialAmount, _name, 10, _symbol)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	address := fmt.Sprintf("0x%x", _tokenAddress)

	// save token into db
	var _tokens bean.Tokens
	_tokens.ESN = "fans"
	_tokens.Address = address
	_tokens.Name = _name
	_tokens.Symbol = _symbol
	_tokens.Total = _int64
	_tokens.Desc = _desc
	_tokens.UserID = _userID
	_tokens.Owner = _addresses[0].Address
	_tokens.Weight = 0
	_tokens.Status = 0

	err = persist.GetPersist().CreateToken(_tokens)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "create token error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": address})
}

// SelfDeployToken SelfDeployToken
func SelfDeployToken(c *gin.Context) {

	_userID := c.PostForm("userID")
	_pwd := c.PostForm("pwd")

	_name := c.PostForm("name")
	_symbol := c.PostForm("symbol")
	_total := c.PostForm("total")
	_desc := c.PostForm("desc")

	_int64, err := strconv.ParseInt(_total, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_addresses, err := persist.GetPersist().ListAddress(_userID)
	if err != nil || len(_addresses) == 0 {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "get user address error"})
		return
	}

	fmt.Println(_addresses[0].Address)

	_keystore, err := persist.GetPersist().AddressInfo(_addresses[0].Address)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "get address error"})
		return
	}

	txOpt, err := bind.NewTransactor(strings.NewReader(_keystore.KeyStore), _pwd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_initialAmount := big.NewInt(_int64)
	_tokenAddress, _, _, err := token.DeployHumanStandardToken(txOpt, ether.GetEthClient(), _initialAmount, _name, 10, _symbol)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	address := fmt.Sprintf("0x%x", _tokenAddress)

	// save token into db
	var _tokens bean.Tokens
	_tokens.ESN = "fans"
	_tokens.Address = address
	_tokens.Name = _name
	_tokens.Symbol = _symbol
	_tokens.Total = _int64
	_tokens.Desc = _desc
	_tokens.UserID = _userID
	_tokens.Owner = _addresses[0].Address
	_tokens.Weight = 0
	_tokens.Status = 0

	err = persist.GetPersist().CreateToken(_tokens)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "create token error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": address})
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
			fmt.Println("Caller Error : " + err.Error())
			continue
		}
		_bigint, err := _caller.BalanceOf(&bind.CallOpts{Pending: false}, common.HexToAddress(_addr))
		if err != nil {
			fmt.Println("BalanceOf Error : " + err.Error())
			continue
		}

		var tokenValue bean.TokenValue
		tokenValue.Address = _conaddr
		tokenValue.Balance = _bigint.String()

		tokenValues = append(tokenValues, tokenValue)
	}

	c.JSON(http.StatusOK, tokenValues)
}

// TransferToken TransferToken
func TransferToken(c *gin.Context) {

	_conaddr := c.PostForm("conaddr")
	_from := c.PostForm("from")
	_to := c.PostForm("to")
	_amount := c.PostForm("amount")
	_pwd := c.PostForm("pwd")
	_memo := c.PostForm("memo")

	_int64, err := strconv.ParseInt(_amount, 10, 64)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	_caller, err := token.NewHumanStandardTokenCaller(common.HexToAddress(_conaddr), ether.GetEthClient())
	if err != nil {
		fmt.Println("Caller Error : " + err.Error())
	}
	_nBalance, err := _caller.BalanceOf(&bind.CallOpts{Pending: false}, common.HexToAddress(_from))
	if err != nil {
		fmt.Println("BalanceOf Error : " + err.Error())
	}

	fmt.Println(_nBalance)
	fmt.Println(_int64)
	if _nBalance.Uint64() < uint64(_int64) {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "余额不足"})
		return
	}

	_keystore, err := persist.GetPersist().AddressInfo(_from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "get address error"})
		return
	}

	txOpt, err := bind.NewTransactor(strings.NewReader(_keystore.KeyStore), _pwd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	ts, err := token.NewHumanStandardTokenTransactor(common.HexToAddress(_conaddr), ether.GetEthClient())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_bigint := big.NewInt(_int64)
	_, err = ts.Transfer(txOpt, common.HexToAddress(_to), _bigint)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	// save transfer into db
	var _transfer bean.TokenTransfer
	_transfer.Address = _conaddr
	_transfer.Type = 0
	_transfer.FromAddr = _from
	_transfer.ToAddr = _to
	_transfer.Amount = _int64
	_transfer.Memo = _memo

	err = persist.GetPersist().CreateTokenTransfer(_transfer)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "create token error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"conaddr": _conaddr,
		"from":    _from,
		"to":      _to,
		"amount":  _amount,
	})
}

// CreateToken CreateToken
func CreateToken(c *gin.Context) {

	var _tokens bean.Tokens
	c.Bind(&_tokens)

	err := persist.GetPersist().CreateToken(_tokens)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "create token error"})
		return
	}

	c.JSON(http.StatusOK, _tokens)
}

// TokenInfo TokenInfo
func TokenInfo(c *gin.Context) {

	_address := c.Params.ByName("address")

	token, err := persist.GetPersist().TokenInfo(_address)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": token})
}

// UpdateTokenWeight UpdateTokenWeight
func UpdateTokenWeight(c *gin.Context) {
	_address := c.Params.ByName("address")
	_weight := c.Params.ByName("weight")
	if _address == "" || _weight == "" {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "There are some empty fields."})
		return
	}

	_iWeight, err := strconv.Atoi(_weight)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	if err := persist.GetPersist().UpdateTokenWeight(_address, _iWeight); err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "update weight error."})
	} else {
		c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": "success"})
	}
}

// ListToken ListToken
func ListToken(c *gin.Context) {
	_search := c.Params.ByName("search")
	_page := c.Params.ByName("page")
	_nPage, _ := strconv.Atoi(_page)

	fmt.Println(_search)
	fmt.Println(_page)

	_tokens, err := persist.GetPersist().ListToken(_search, _nPage)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _tokens})
}

// CountToken CountToken
func CountToken(c *gin.Context) {

	_count, err := persist.GetPersist().CountToken()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _count})
}

// CreateTokenTransfer CreateTokenTransfer
func CreateTokenTransfer(c *gin.Context) {

	var _transfer bean.TokenTransfer
	c.Bind(&_transfer)

	err := persist.GetPersist().CreateTokenTransfer(_transfer)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "create token error"})
		return
	}

	c.JSON(http.StatusOK, _transfer)
}

// ListTokenTransfer ListTokenTransfer
func ListTokenTransfer(c *gin.Context) {

	_tokenaddress := c.Params.ByName("tokenaddress")
	_address := c.Params.ByName("address")
	_page := c.Params.ByName("page")
	_nPage, _ := strconv.Atoi(_page)

	_transfers, err := persist.GetPersist().ListTokenTransfer(_tokenaddress, _address, _nPage)
	if err != nil {
		c.JSON(406, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _transfers})
}

// AllTokenTransfer AllTokenTransfer
func AllTokenTransfer(c *gin.Context) {

	_tokenaddress := c.Params.ByName("tokenaddress")
	_page := c.Params.ByName("page")
	_nPage, _ := strconv.Atoi(_page)

	_transfers, err := persist.GetPersist().AllTokenTransfer(_tokenaddress, _nPage)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _transfers})
}

// CountTokenTransfer CountTokenTransfer
func CountTokenTransfer(c *gin.Context) {

	_tokenaddress := c.Params.ByName("tokenaddress")

	_count, err := persist.GetPersist().CountTokenTransfer(_tokenaddress)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _count})
}

// CountAllTokenTransfer CountAllTokenTransfer
func CountAllTokenTransfer(c *gin.Context) {

	_count, err := persist.GetPersist().CountAllTokenTransfer()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _count})
}

// SumAllTokenTransfer SumAllTokenTransfer
func SumAllTokenTransfer(c *gin.Context) {

	_count, err := persist.GetPersist().SumAllTokenTransfer()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _count})
}
