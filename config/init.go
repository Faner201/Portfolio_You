package config

import "github.com/spf13/viper"

func Init() error {
	viper.AddConfigPath("/Users/fanfurick/Documents/Profile_You/config/")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
