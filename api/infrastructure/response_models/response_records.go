package response_models

type ResponseRecord struct {
	ID              string
	MobileNumber    string
	Status          string
	ConversationID  string
	MessagesCounter int
	CreatedAt       string
}
