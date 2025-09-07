package ws

import (
	twilio "AISale/services/twillio"
	"log"
	"time"
)

func PollTwilio(chatID, accountSID, authToken string) {
	var lastMessageSID string

	for {
		time.Sleep(5 * time.Second)

		messages, err := fetchMessagesFromTwilio(chatID, lastMessageSID, accountSID, authToken)
		if err != nil {
			log.Println("Twilio fetch error:", err)
			continue
		}

		for _, m := range messages {
			Broadcast(chatID, m.Body)
			lastMessageSID = m.Sid
		}
	}
}

func fetchMessagesFromTwilio(chatID, lastSID, accountSID, authToken string) ([]twilio.Message, error) {
	botNumber := "+16693420294"

	twilioClient := twilio.NewClient(accountSID, authToken)

	messages, err := twilioClient.GetConversation(chatID, botNumber, 50)
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
