package config

import (
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("db.url", "postgres://user:password@localhost/gobot?sslmode=disable")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func GetString(key string) string {
	return viper.GetString(key)
}
