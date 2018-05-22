package persist

import (
	"github.com/Eric-GreenComb/x-server/bean"
)

// AdminLogin AdminLogin Persist
func (persist *Persist) AdminLogin(userID, password string) (bean.AdminUsers, error) {
	var user bean.AdminUsers
	err := persist.db.Where("user_id = ? AND passwd = ?", userID, password).First(&user).Error
	return user, err
}

// CreateAdminUser CreateAdminUser Persist
func (persist *Persist) CreateAdminUser(user bean.AdminUsers) error {
	err := persist.db.Create(&user).Error
	return err
}

// AdminUserInfo AdminUserInfo Persist
func (persist *Persist) AdminUserInfo(userID string) (bean.AdminUsers, error) {

	var user bean.AdminUsers
	err := persist.db.Table("admin_users").Where("user_id = ?", userID).First(&user).Error

	return user, err
}

// UpdateAdminUserPasswd UpdateAdminUserPasswd Persist
func (persist *Persist) UpdateAdminUserPasswd(userID, new string) error {
	return persist.db.Table("admin_users").Where("user_id = ?", userID).Update("passwd", new).Error
}
