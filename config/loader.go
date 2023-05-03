package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AutomaticEnv()
	viper.SetConfigName("application")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./../")
	viper.AddConfigPath("./../../")
	fmt.Println(viper.ReadInConfig())
	initDbConfig()
}
