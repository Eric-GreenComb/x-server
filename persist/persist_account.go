package persist

import (
	"github.com/Eric-GreenComb/x-server/bean"
)

// CreateAddress CreateAddress Persist
func (persist *Persist) CreateAddress(address bean.Addresses) error {
	err := persist.db.Create(&address).Error
	return err
}

// AddressInfo AddressInfo Persist
func (persist *Persist) AddressInfo(addr string) (bean.Addresses, error) {

	var address bean.Addresses
	err := persist.db.Table("addresses").Where("address = ?", addr).First(&address).Error

	return address, err
}

// ListAddress ListAddress Persist
func (persist *Persist) ListAddress(userID string) ([]bean.Addresses, error) {

	var addresses []bean.Addresses
	err := persist.db.Table("addresses").Select("created_at, address, user_id, name").Where("user_id = ?", userID).Find(&addresses).Error

	return addresses, err
}

// UpdateAccountPwd UpdateAccountPwd Persist
func (persist *Persist) UpdateAccountPwd(userID, address, keystore string) error {
	return persist.db.Table("addresses").Where("user_id = ? and address = ?", userID, address).Update("key_store", keystore).Error
}
