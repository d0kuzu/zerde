package response_models

type ResponseRecord struct {
	ID              string `json:"id"`
	MobileNumber    string `json:"mobileNumber"`
	Status          string `json:"status"`
	ConversationID  string `json:"conversationId"`
	MessagesCounter int    `json:"messagesCounter"`
	CreatedAt       string `json:"createdTime"`
}
