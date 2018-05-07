package persist

import (
	"github.com/Eric-GreenComb/x-server/bean"
)

// Login Login Persist
func (persist *Persist) Login(userID, password string) (bean.Users, error) {
	var user bean.Users
	err := persist.db.Where("user_id = ? AND passwd = ?", userID, password).First(&user).Error
	return user, err
}

// CreateUser CreateUser Persist
func (persist *Persist) CreateUser(user bean.Users) error {
	err := persist.db.Create(&user).Error
	return err
}

// UserInfo UserInfo Persist
func (persist *Persist) UserInfo(userID string) (bean.Users, error) {

	var user bean.Users
	err := persist.db.Table("users").Where("user_id = ?", userID).First(&user).Error

	return user, err
}

// UpdateUserPasswd UpdateUserPasswd Persist
func (persist *Persist) UpdateUserPasswd(userID, new string) error {
	return persist.db.Table("users").Where("user_id = ?", userID).Update("passwd", new).Error
}
