package bean

import ()

// TokenValue TokenValue
type TokenValue struct {
	Address string `form:"address" json:"address"`               // token address
	Name    string `gorm:"not null" form:"name" json:"name"`     // Token合约名称
	Symbol  string `gorm:"not null" form:"symbol" json:"symbol"` // Token合约symbol
	Balance string `form:"balance" json:"balance"`               // balance
}
