package config

import (
	"github.com/spf13/viper"
)

// Config struct to hold configuration values
type Config struct {
	Port               string
	UserBaseURL        string
	TransactionBaseURL string
}

// InitConfig initializes and reads the configuration values from .env file
func InitConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{
		Port:               viper.GetString("PORT"),
		UserBaseURL:        viper.GetString("USER_BASE_URL"),
		TransactionBaseURL: viper.GetString("TRANSACTION_BASE_URL"),
	}

	return config, nil
}
