package twilio

type Message struct {
	Sid         string `json:"sid"`
	From        string `json:"from"`
	To          string `json:"to"`
	Body        string `json:"body"`
	DateSent    string `json:"date_sent"`
	DateCreated string `json:"date_created"`
}

type MessagesResponse struct {
	Messages []Message `json:"messages"`
}
