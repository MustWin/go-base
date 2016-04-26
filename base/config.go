package base

import (
	"github.com/spf13/viper"
)

func loadConfig() {
	viper.SetConfigName("config")
	viper.ReadInConfig()
}
