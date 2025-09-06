package ws

import (
	twilio "AISale/services/twillio"
	"fmt"
	"log"
	"sort"
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
			fmt.Printf(lastMessageSID)
		}
	}
}

func fetchMessagesFromTwilio(chatID, lastSID, accountSID, authToken string) ([]twilio.Message, error) {
	botNumber := "+16693420294"
	twilioClient := twilio.NewClient(accountSID, authToken)

	allMessages, err := twilioClient.GetConversation(chatID, botNumber, 50)
	if err != nil {
		return nil, err
	}

	if lastSID == "" {
		sort.Slice(allMessages, func(i, j int) bool {
			return allMessages[i].DateCreated < allMessages[j].DateCreated
		})
		return allMessages, nil
	}

	var result []twilio.Message
	for _, msg := range allMessages {
		if msg.Sid == lastSID {
			break
		}
		result = append(result, msg)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].DateCreated < result[j].DateCreated
	})

	return result, nil
}
