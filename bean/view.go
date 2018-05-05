package bean

import ()

// TokenValue TokenValue
type TokenValue struct {
	Address string `form:"address" json:"address"` // token address
	Balance string `form:"balance" json:"balance"` // balance
}
