package ws

import (
	"AISale/config"
	twilio "AISale/services/twillio"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func PollTwilio(chatID, accountSID, authToken string) {
	var lastMessageSID string

	for {
		messages, err := fetchMessagesFromTwilio(chatID, lastMessageSID, accountSID, authToken)
		if err != nil {
			log.Println("Twilio fetch error:", err)
			continue
		}

		for _, m := range messages {
			fmt.Println(m.Body)
			var author string
			if m.From != config.BotNumber {
				author = "bot"
			} else {
				author = "client"
			}

			msg := Message{
				Author: author,
				Body:   m.Body,
			}

			data, err := json.Marshal(msg)
			if err != nil {
				log.Println("ws message json marshal error:", err)
			}

			Broadcast(chatID, data)

			lastMessageSID = m.Sid
		}
		fmt.Println("sended")

		time.Sleep(3 * time.Second)
	}
}

func fetchMessagesFromTwilio(chatID, lastSID, accountSID, authToken string) ([]twilio.Message, error) {

	twilioClient := twilio.NewClient(accountSID, authToken)

	messages, err := twilioClient.GetConversation(chatID, config.BotNumber, 1000)
	if err != nil {
		return nil, err
	}

	if lastSID == "" {
		return messages, nil
	}

	var result []twilio.Message
	found := false
	for _, msg := range messages {
		if found {
			result = append(result, msg)
		}
		if msg.Sid == lastSID {
			found = true
		}
	}

	return result, nil
}
