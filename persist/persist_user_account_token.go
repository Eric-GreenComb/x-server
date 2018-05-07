package persist

import (
	"github.com/Eric-GreenComb/x-server/bean"
)

// CreateUserAddressTokens CreateUserAddressTokens Persist
func (persist *Persist) CreateUserAddressTokens(userAddressToken bean.UserAddressTokens) error {
	err := persist.db.Create(&userAddressToken).Error
	return err
}

// DeleteUserAddressTokens DeleteUserAddressTokens Persist
func (persist *Persist) DeleteUserAddressTokens(userID, address, tokenAddress string) error {

	var userAddressToken bean.UserAddressTokens
	err := persist.db.Table("user_address_tokens").Where("user_id = ? and address = ? and token_address = ?", userID, address, tokenAddress).Delete(&userAddressToken).Error
	return err
}

// GetUserAddressTokens GetUserAddressTokens Persist
func (persist *Persist) GetUserAddressTokens(userID, address, tokenAddress string) (bean.UserAddressTokens, error) {

	var userAddressToken bean.UserAddressTokens
	err := persist.db.Table("user_address_tokens").Where("user_id = ? and address = ? and token_address = ?", userID, address, tokenAddress).First(&userAddressToken).Error

	return userAddressToken, err
}

// ListUserAddressTokens ListUserAddressTokens Persist
func (persist *Persist) ListUserAddressTokens(userID, address string) ([]bean.UserAddressTokens, error) {

	var userAddressTokens []bean.UserAddressTokens
	err := persist.db.Table("user_address_tokens").Where("user_id = ? and address = ?", userID, address).Find(&userAddressTokens).Error

	return userAddressTokens, err
}
