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
func (persist *Persist) ListToken(search string, page int) ([]bean.Tokens, error) {

	var tokens []bean.Tokens

	page--
	if page < 0 {
		page = 0
	}

	_0ffset := page * config.ServerConfig.ViewLimit

	var err error

	if len(search) == 0 {
		err = persist.db.Table("tokens").Order("weight desc").Limit(config.ServerConfig.ViewLimit).Offset(_0ffset).Find(&tokens).Error
		return tokens, nil
	}

	err = persist.db.Table("tokens").Where("user_id like ?", "%"+search+"%").Order("weight desc").Limit(config.ServerConfig.ViewLimit).Offset(_0ffset).Find(&tokens).Error
	return tokens, err
}

// CountToken CountToken
func (persist *Persist) CountToken() (uint64, error) {

	var value uint64
	err := persist.db.Table("tokens").Count(&value).Error

	if err != nil {
		return 0, err
	}
	return value, nil

}

// CreateTokenTransfer CreateTokenTransfer Persist
func (persist *Persist) CreateTokenTransfer(transfer bean.TokenTransfer) error {
	err := persist.db.Create(&transfer).Error
	return err
}

// ListTokenTransfer ListTokenTransfer
func (persist *Persist) ListTokenTransfer(tokenAddr, addr string, page int) ([]bean.TokenTransfer, error) {

	var transfers []bean.TokenTransfer

	page--
	if page < 0 {
		page = 0
	}

	_0ffset := page * config.ServerConfig.ViewLimit

	var err error

	err = persist.db.Table("token_transfers").Where("address = ? AND (from_addr = ? OR to_addr = ?)", tokenAddr, addr, addr).Order("id desc").Limit(config.ServerConfig.ViewLimit).Offset(_0ffset).Find(&transfers).Error

	if err != nil {
		return nil, err
	}
	return transfers, nil
}

// AllTokenTransfer AllTokenTransfer
func (persist *Persist) AllTokenTransfer(tokenAddr string, page int) ([]bean.TokenTransfer, error) {

	var transfers []bean.TokenTransfer

	page--
	if page < 0 {
		page = 0
	}

	_0ffset := page * config.ServerConfig.ViewLimit

	var err error

	err = persist.db.Table("token_transfers").Where("address = ?", tokenAddr).Order("id desc").Limit(config.ServerConfig.ViewLimit).Offset(_0ffset).Find(&transfers).Error

	if err != nil {
		return nil, err
	}
	return transfers, nil
}

// CountTokenTransfer CountTokenTransfer
func (persist *Persist) CountTokenTransfer(tokenAddr string) (uint64, error) {

	var value uint64
	err := persist.db.Table("token_transfers").Where("address = ?", tokenAddr).Count(&value).Error

	if err != nil {
		return 0, err
	}
	return value, nil
}

// CountAllTokenTransfer CountAllTokenTransfer
func (persist *Persist) CountAllTokenTransfer() (uint64, error) {

	var value uint64
	err := persist.db.Table("token_transfers").Count(&value).Error

	if err != nil {
		return 0, err
	}
	return value, nil
}

// Result Result
type Result struct {
	Total int
}

// SumAllTokenTransfer SumAllTokenTransfer
func (persist *Persist) SumAllTokenTransfer() (int, error) {

	var result Result

	err := persist.db.Table("token_transfers").Select("sum(amount) as total").Scan(&result).Error

	if err != nil {
		return 0, err
	}
	return result.Total, nil
}
