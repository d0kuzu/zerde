package airtable

import "AISale/api/infrastructure/response_models"

//type Record struct {
//	ID     string                 `json:"id"`
//	Fields map[string]interface{} `json:"fields"`
//}

type PurchaseFields struct {
	MobileNumber    string `json:"Mobile Number,omitempty"`
	Status          string `json:"Status,omitempty"`
	ConversationID  string `json:"Conversation ID,omitempty"`
	MessagesCounter int    `json:"-"`
}

type Record struct {
	ID        string         `json:"id"`
	Fields    PurchaseFields `json:"fields"`
	CreatedAt string         `json:"createdTime"`
}

func (r *Record) ToResponse() response_models.ResponseRecord {
	return response_models.ResponseRecord{
		ID:              r.ID,
		MobileNumber:    r.Fields.MobileNumber,
		Status:          r.Fields.Status,
		ConversationID:  r.Fields.ConversationID,
		MessagesCounter: r.Fields.MessagesCounter,
		CreatedAt:       r.CreatedAt,
	}
}

type ListResponse struct {
	Records []Record `json:"records"`
	Offset  string   `json:"offset,omitempty"`
}
