package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/x-server/badger"
)

// SetBadgerKey SetBadgerKey
func SetBadgerKey(c *gin.Context) {

	_key := c.Params.ByName("key")
	_value := c.Params.ByName("value")

	badger.NewWrite().Set(_key, []byte(_value))

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "key": _key, "value": _value})
}

// SetBadgerKeyTTL SetBadgerKeyTTL
func SetBadgerKeyTTL(c *gin.Context) {

	_key := c.Params.ByName("key")
	_value := c.Params.ByName("value")
	_ttlString := c.Params.ByName("ttl")

	_nTTL, _ := strconv.Atoi(_ttlString)
	second := int(time.Second)

	badger.NewWrite().SetWithTTL(_key, []byte(_value), time.Duration(_nTTL*second))

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "key": _key, "value": _value, "ttl": _ttlString})
}

// GetBadgerKey GetBadgerKey
func GetBadgerKey(c *gin.Context) {

	_key := c.Params.ByName("key")

	_value, err := badger.NewRead().Get(_key)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "errinfo": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "key": _key, "value": string(_value)})
}
