package chat

import (
	"AISale/config"
	"AISale/services/GPT"
	"AISale/services/qdrant"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"log"
	"strings"
)

func Conversation(c *gin.Context, userId string, userMessage string, consType config.ConservationType) (string, error) {
	log.Printf("Сообщение от пользователя %s: %s", userId, userMessage)
	messages, err := GetMessages(userId, consType)
	if err != nil {
		return "", err
	}

	if consType == config.Remind {
		AddMessage(&messages, "system", userMessage)
	} else {
		AddMessage(&messages, "user", userMessage)
	}

	response, err := GPT.GetAnswer(c, messages, consType)
	if err != nil {
		return "", err
	}
	log.Printf("Ответ пользователю %s от ИИ: %s\n", userId, response.Choices[0].Message.Content)

	jsonText := GetJSONFromText(response.Choices[0].Message.Content)
	if len(jsonText) != 0 {
		log.Printf("Данные запроса пользователя %s: %s\n", userId, jsonText)

		var queriesAnswers []string

		for _, query := range jsonText {
			answers, err := qdrant.SearchVector(c.Request.Context(), "my_collection", query, 3)
			if err != nil {
				return "", err
			}
			joinedAnswers := fmt.Sprintf("Ответы по запросу '%s' : \n%s", query, strings.Join(answers, "\n"))
			queriesAnswers = append(queriesAnswers, joinedAnswers)

			log.Printf("Ответы от qdrant: %s\n", joinedAnswers)
		}
		joinedAnswers := strings.Join(queriesAnswers, "\n\n")
		prompt := fmt.Sprintf(config.QdrantAnswerPrompt, joinedAnswers)
		promptMessage := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: prompt,
		}
		response, err = GPT.GetAnswer(c, []openai.ChatCompletionMessage{promptMessage}, consType)
		if err != nil {
			return "", err
		}
	}

	AddMessage(&messages, "assistant", response.Choices[0].Message.Content)

	err = SaveMessages(userId, messages)
	if err != nil {
		return "", err
	}

	log.Println("Конец запроса\n\n\n")
	return response.Choices[0].Message.Content, nil
}
