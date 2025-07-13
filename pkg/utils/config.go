package utils

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	AppName    string
	Port       string
	Debug      bool
	DB         DatabaseConfig
	Limit      int
	PathLogger string
	PathUpload string
	JWTKey     string
}

type DatabaseConfig struct {
	Name     string
	Username string
	Password string
	Host     string
	Port     string
}

func ReadConfiguration() (Configuration, error) {
	viper.SetConfigFile(".env") // read file .env
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		return Configuration{}, err
	}
	viper.AutomaticEnv() // read env os

	return Configuration{
		AppName:    viper.GetString("APP_NAME"),
		Port:       viper.GetString("PORT"),
		Debug:      viper.GetBool("DEBUG"),
		Limit:      viper.GetInt("LIMIT"),
		PathLogger: viper.GetString("PATH_LOGGER"),
		PathUpload: viper.GetString("PATH_UPLOAD"),
		JWTKey:     viper.GetString("JWT_KEY"),
		DB: DatabaseConfig{
			Name:     viper.GetString("DATABASE_NAME"),
			Username: viper.GetString("DATABASE_USER"),
			Password: viper.GetString("DATABASE_PASSWORD"),
			Host:     viper.GetString("DATABASE_HOST"),
			Port:     viper.GetString("SHELL"),
		},
	}, nil
}
