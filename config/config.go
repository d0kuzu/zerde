package config

import (
	"AISale/services/chrome"
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	Chrome *chrome.Client
	Cfg    *Settings
}

func NewApp(chrome *chrome.Client, cfg *Settings) *App {
	return &App{
		Chrome: chrome,
		Cfg:    cfg,
	}
}

type Settings struct {
	ApiKey    string
	BaseID    string
	TableName string

	AccountSID string
	AuthToken  string

	DiaxelLogin    string
	DiaxelPassword string

	ChroniumPath string
}

func LoadConfig() (*Settings, error) {
	godotenv.Load(".env")

	return &Settings{
		ApiKey:         os.Getenv("API_KEY"),
		BaseID:         os.Getenv("BASE_ID"),
		TableName:      os.Getenv("TABLE_NAME"),
		AccountSID:     os.Getenv("ACCOUNT_SID"),
		AuthToken:      os.Getenv("AUTH_TOKEN"),
		DiaxelLogin:    os.Getenv("DIAXEL_LOGIN"),
		DiaxelPassword: os.Getenv("DIAXEL_PASSWORD"),
		ChroniumPath:   os.Getenv("CHROMIUM_PATH"),
	}, nil
}
