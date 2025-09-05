package config

import (
	"github.com/spf13/viper"
)

type Settings struct {
	ApiKey    string
	BaseID    string
	TableName string

	AccountSID string
	AuthToken  string
}

func LoadConfig() (*Settings, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Settings{
		ApiKey:    viper.GetString("API_KEY"),
		BaseID:    viper.GetString("BASE_ID"),
		TableName: viper.GetString("TABLE_NAME"),

		AccountSID: viper.GetString("ACCOUNT_SID"),
		AuthToken:  viper.GetString("AUTH_TOKEN"),
	}, nil
}
