package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Settings struct {
	ApiKey    string
	BaseID    string
	TableName string

	AccountSID string
	AuthToken  string
}

func LoadConfig() (*Settings, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &Settings{
		ApiKey:     os.Getenv("API_KEY"),
		BaseID:     os.Getenv("BASE_ID"),
		TableName:  os.Getenv("TABLE_NAME"),
		AccountSID: os.Getenv("ACCOUNT_SID"),
		AuthToken:  os.Getenv("AUTH_TOKEN"),
	}, nil
}
