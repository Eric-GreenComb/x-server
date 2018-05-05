package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"

	"github.com/Eric-GreenComb/x-server/bean"
)

// DBConfig 数据库相关配置
var DBConfig bean.DBConfig

// ServerConfig Server Config
var ServerConfig bean.ServerConfig

func init() {
	readConfig()
	initConfig()
}

func readConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
}

func initConfig() {
	ServerConfig.Port = strings.Split(viper.GetString("server.port"), ",")
	ServerConfig.Mode = viper.GetString("server.mode")
	ServerConfig.GormLogMode = viper.GetString("server.gorm.LogMode")
	ServerConfig.ViewLimit = viper.GetInt("server.view.limit")

	DBConfig.Dialect = viper.GetString("database.dialect")
	DBConfig.Database = viper.GetString("database.database")
	DBConfig.User = viper.GetString("database.user")
	DBConfig.Password = viper.GetString("database.password")
	DBConfig.Host = viper.GetString("database.host")
	DBConfig.Port = viper.GetInt("database.port")
	DBConfig.Charset = viper.GetString("database.charset")
	DBConfig.MaxIdleConns = viper.GetInt("database.maxIdleConns")
	DBConfig.MaxOpenConns = viper.GetInt("database.maxOpenConns")
	DBConfig.URL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		DBConfig.User, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Database, DBConfig.Charset)
}
