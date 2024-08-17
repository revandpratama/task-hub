package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PORT           string
	JWT_SECRET_KEY string
	DB_URL         string
	DB_PORT        string
	DB_NAME        string
	DB_USERNAME    string
	DB_PASSWORD    string
}

var ENV *Config

func LoadConfig() {

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("config not loaded properly: %v", err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		log.Fatalf("unable to unmarshal config: %v", err)
	}

	log.Println("config loaded...")

}
