package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

func GetString(key string) string {
	value := os.Getenv(key)
	if strings.TrimSpace(value) == "" {
		value = viper.GetString(key)
	}

	if strings.TrimSpace(value) == "" {
		panic(fmt.Sprintf("Key %s is not set", key))
	}

	return value
}

func GetInt(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err == nil {
		return value
	}

	value, err = strconv.Atoi(viper.GetString(key))
	if err == nil {
		return value
	}

	panic(fmt.Sprintf("Key: %s invalid value set", key))
}
