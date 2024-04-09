package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort     string `mapstructure:"PORT"`
	ServiceName string `mapstructure:"SERVICE_NAME"`
}

func Initialize() (config Config) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error to read .env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error to load enviroment variables")
	}
	return
}
