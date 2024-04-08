package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort     string
	ServiceName string
}

func Initialize() (config Config) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error to read .env file", err)
	}

	config.AppPort = viper.GetString("PORT")
	config.ServiceName = viper.GetString("SERVICE_NAME")
	return
}
