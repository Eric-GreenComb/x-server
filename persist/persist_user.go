package persist

import (
	"github.com/Eric-GreenComb/x-server/bean"
	"github.com/Eric-GreenComb/x-server/config"
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

// CountUser CountUser
func (persist *Persist) CountUser() (uint64, error) {

	var value uint64
	err := persist.db.Table("users").Count(&value).Error

	if err != nil {
		return 0, err
	}
	return value, nil
}

// ListUser ListUser Persist
func (persist *Persist) ListUser(search string, page int) ([]bean.Users, error) {

	var users []bean.Users

	page--
	if page < 0 {
		page = 0
	}

	_0ffset := page * config.ServerConfig.ViewLimit

	var err error

	if len(search) == 0 {
		err = persist.db.Table("users").Select("*").Limit(config.ServerConfig.ViewLimit).Offset(_0ffset).Find(&users).Error
		return users, err
	}

	err = persist.db.Table("users").Where("user_id like ?", "%"+search+"%").Select("*").Limit(config.ServerConfig.ViewLimit).Offset(_0ffset).Find(&users).Error

	return users, err
}
