package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	DbDriver   string `mapstructure:"DB_DRIVER"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbName     string `mapstructure:"DB_Name"`
	Domain     string `mapstructure:"DOMAIN"`
}

var AppConfig Config

func LoadConfig() error {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
