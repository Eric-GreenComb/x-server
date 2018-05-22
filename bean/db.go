package bean

import (
	"github.com/jinzhu/gorm"
)

// Ecologic 生态
type Ecologic struct {
	gorm.Model
	ESN   string  `gorm:"not null" form:"esn" json:"esn"`     // 生态SN
	EName string  `gorm:"not null" form:"ename" json:"ename"` // 生态名称
	Owner string  `gorm:"not null" form:"owner" json:"owner"` // 生态管理者
	Desc  string  `gorm:"not null" form:"desc" json:"desc"`   // 生态描述
	Nodes []Nodes `form:"nodes" json:"nodes"`                 // 节点信息
}

// Nodes Ecologic nodes
type Nodes struct {
	ESN    string `gorm:"not null" form:"esn" json:"esn"`       // 生态SN
	RawURL string `gorm:"not null" form:"rawurl" json:"rawurl"` // 节点链接
}

// Tokens Tokens
type Tokens struct {
	gorm.Model
	ESN     string `gorm:"not null" form:"esn" json:"esn"`         // 生态SN
	Address string `gorm:"not null" form:"address" json:"address"` // Token合约地址
	Name    string `gorm:"not null" form:"name" json:"name"`       // Token合约名称
	Symbol  string `gorm:"not null" form:"symbol" json:"symbol"`   // Token合约symbol
	Total   int64  `gorm:"not null" form:"total" json:"total"`     // Token合约total
	Desc    string `gorm:"not null" form:"desc" json:"desc"`       // Token合约描述
	UserID  string `gorm:"not null" form:"userID" json:"userID"`   // Token合约发行者
	Owner   string `gorm:"not null" form:"owner" json:"owner"`     // Token合约发行者地址
	Weight  int    `gorm:"not null" form:"weight" json:"weight"`   // Token合约权重
	Status  int8   `gorm:"not null" form:"status" json:"status"`   // Token合约状态
}

// TokenTransfer TokenTransfer
type TokenTransfer struct {
	gorm.Model
	Address  string `gorm:"not null" form:"address" json:"address"` // Token合约地址
	Type     int8   `gorm:"not null" form:"type" json:"type"`       // loan,transfer,recharge,withdraw
	FromAddr string `gorm:"default:''" form:"from" json:"from"`
	ToAddr   string `gorm:"default:''" form:"to" json:"to"`
	Amount   int64  `gorm:"not null" form:"amount" json:"amount"`
	Memo     string `gorm:"default:''" form:"memo" json:"memo"` // coinbase,mint,transfer,recharge,withdraw
}

// Users Users
type Users struct {
	gorm.Model
	User
}

// User User
type User struct {
	UserID string `gorm:";unique" form:"userID" json:"userID"` // 手机号码
	Name   string `gorm:"not null" form:"name" json:"name"`
	Passwd string `gorm:"not null" form:"passwd" json:"passwd"`
	Email  string `gorm:"not null" form:"email" json:"email"`
}

// AdminUsers AdminUsers
type AdminUsers struct {
	gorm.Model
	User
}

// UserAddressTokens UserAddressTokens
type UserAddressTokens struct {
	gorm.Model
	UserID       string `gorm:"not null" form:"userID" json:"userID"`             // 手机号码
	Address      string `gorm:"not null" form:"address" json:"address"`           // 用户地址
	ESN          string `gorm:"not null" form:"esn" json:"esn"`                   // 生态SN
	TokenAddress string `gorm:"not null" form:"tokenaddress" json:"tokenaddress"` // Token合约地址
	Name         string `gorm:"not null" form:"name" json:"name"`                 // Token合约名称
	Symbol       string `gorm:"not null" form:"symbol" json:"symbol"`             // Token合约symbol
}

// Addresses Addresses
type Addresses struct {
	gorm.Model
	Address   string `gorm:";unique" form:"address" json:"address"` // 用户地址
	UserID    string `gorm:"not null" form:"userID" json:"userID"`  // 手机号码
	Name      string `gorm:"not null" form:"name" json:"name"`
	KeyStore  string `gorm:"size:600;not null" form:"keystore" json:"keystore"`   // 用户keystore
	BackStore string `gorm:"size:600;not null" form:"backstore" json:"backstore"` // 备份keystore
}
