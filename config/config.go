package config

import (
	"LoginUser/utils"
	"fmt"
)

type AppConfig struct {
	Username        string
	Password        string
	Host            string
	Database        string
	SecretKeyLength int
}

var appConfig AppConfig

func initDbConfig() {
	appConfig = AppConfig{
		Username:        utils.GetString("MYSQL_USERNAME"),
		Password:        utils.GetString("MYSQL_PASSWORD"),
		Host:            utils.GetString("MYSQL_HOST"),
		Database:        utils.GetString("MYSQL_DATABASE"),
		SecretKeyLength: utils.GetInt("SECRET_LENGTH"),
	}
}

func GetAppConfig() AppConfig {
	return appConfig
}

func (a AppConfig) GetAppConfigUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", a.Username, a.Password, a.Host, a.Database)
}
