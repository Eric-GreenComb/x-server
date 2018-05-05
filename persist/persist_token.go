package persist

import (
	"github.com/Eric-GreenComb/x-server/bean"
	"github.com/Eric-GreenComb/x-server/config"
)

// CreateToken CreateToken Persist
func (persist *Persist) CreateToken(token bean.Tokens) error {
	err := persist.db.Create(&token).Error
	return err
}

// TokenInfo TokenInfo Persist
func (persist *Persist) TokenInfo(addr string) (bean.Tokens, error) {

	var token bean.Tokens
	err := persist.db.Table("tokens").Where("address = ?", addr).First(&token).Error

	return token, err
}

// UpdateTokenWeight UpdateTokenWeight Persist
func (persist *Persist) UpdateTokenWeight(addr string, weight int) error {
	return persist.db.Table("tokens").Where("address = ?", addr).Update(map[string]interface{}{"weight": weight}).Error
}

// ListToken ListToken
func (persist *Persist) ListToken(page int) ([]bean.Tokens, error) {

	var tokens []bean.Tokens

	page--
	if page < 0 {
		page = 0
	}

	_0ffset := page * config.ServerConfig.ViewLimit

	var err error

	err = persist.db.Table("tokens").Order("weight desc").Limit(config.ServerConfig.ViewLimit).Offset(_0ffset).Find(&tokens).Error

	if err != nil {
		return nil, err
	}
	return tokens, nil
}

// CreateTokenTransfer CreateTokenTransfer Persist
func (persist *Persist) CreateTokenTransfer(transfer bean.TokenTransfer) error {
	err := persist.db.Create(&transfer).Error
	return err
}

// ListTokenTransfer ListTokenTransfer
func (persist *Persist) ListTokenTransfer(addr string, page int) ([]bean.TokenTransfer, error) {

	var transfers []bean.TokenTransfer

	page--
	if page < 0 {
		page = 0
	}

	_0ffset := page * config.ServerConfig.ViewLimit

	var err error

	err = persist.db.Table("token_transfers").Where("address = ?", addr).Order("id desc").Limit(config.ServerConfig.ViewLimit).Offset(_0ffset).Find(&transfers).Error

	if err != nil {
		return nil, err
	}
	return transfers, nil
}
