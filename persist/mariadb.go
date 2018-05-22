package persist

import (
	"github.com/jinzhu/gorm"

	"github.com/Eric-GreenComb/x-server/bean"
	"github.com/Eric-GreenComb/x-server/config"
)

// ConnectDb connect Db
func ConnectDb() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)

	if config.ServerConfig.GormLogMode == "false" {
		db.LogMode(false)
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}

// InitDatabase Init Db
func InitDatabase() {
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	defer db.Close()

	if config.ServerConfig.GormLogMode == "false" {
		db.LogMode(false)
	}

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&bean.Ecologic{}) {
		db.CreateTable(&bean.Ecologic{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&bean.Ecologic{})
	}

	if !db.HasTable(&bean.Nodes{}) {
		db.CreateTable(&bean.Nodes{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&bean.Nodes{})
	}

	if !db.HasTable(&bean.Tokens{}) {
		db.CreateTable(&bean.Tokens{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&bean.Tokens{})
	}

	if !db.HasTable(&bean.TokenTransfer{}) {
		db.CreateTable(&bean.TokenTransfer{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&bean.TokenTransfer{})
	}

	if !db.HasTable(&bean.Users{}) {
		db.CreateTable(&bean.Users{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&bean.Users{})
	}

	if !db.HasTable(&bean.UserAddressTokens{}) {
		db.CreateTable(&bean.UserAddressTokens{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&bean.UserAddressTokens{})
	}

	if !db.HasTable(&bean.Addresses{}) {
		db.CreateTable(&bean.Addresses{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&bean.Addresses{})
	}

	if !db.HasTable(&bean.AdminUsers{}) {
		db.CreateTable(&bean.AdminUsers{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&bean.AdminUsers{})
	}

	return
}
