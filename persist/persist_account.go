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
	err := persist.db.Table("addresses").Select("address, user_id").Where("user_id = ?", userID).Find(&addresses).Error

	return addresses, err
}
