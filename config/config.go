package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Settings struct {
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	DbPort     string
	Ssl        string

	QdrantGRPCHost string

	OpenaiApiKey string

	TwilioUsername string
	TwilioToken    string
	TwilioNumber   string
}

func LoadENV() (Settings, error) {
	err := godotenv.Load()
	if err != nil {
		return Settings{}, errors.New("error loading .env file: " + err.Error())
	}

	settings := Settings{
		DbHost:     os.Getenv("DB_HOST"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbPort:     os.Getenv("DB_PORT"),
		Ssl:        os.Getenv("DB_SSL"),

		QdrantGRPCHost: os.Getenv("QDRANT_GRPC_HOST"),

		OpenaiApiKey: os.Getenv("OPENAI_API_KEY"),

		TwilioUsername: os.Getenv("TWILIO_ACCOUNT_SID"),
		TwilioToken:    os.Getenv("TWILIO_AUTH_TOKEN"),
		TwilioNumber:   os.Getenv("TWILIO_PHONE_NUMBER"),
	}

	return settings, nil
}

const WaitingTime = time.Minute
const MaxRemindCount = 2
