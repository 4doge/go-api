package config

import "github.com/spf13/viper"

var config *viper.Viper

func Init(env string) {
	config = viper.New()
	config.SetConfigType("json")
	config.SetConfigName(env)
	config.AddConfigPath("config/")
	err := config.ReadInConfig()
	if err != nil {
		panic("Can't read config file")
	}
}

func GetConfig() *viper.Viper {
	return config
}
