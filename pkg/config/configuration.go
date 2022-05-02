package config

import (
	"github.com/spf13/viper"
	"log"
)

var Config *Configuration

type Configuration struct {
	Server  ServerConfiguration
	MongoDB MongoConfiguration
}

// Setup initialize configuration
func Setup() error {
	var configuration *Configuration

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
		return err
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
		return err
	}

	Config = configuration

	return nil
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
