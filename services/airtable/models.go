package airtable

//type Record struct {
//	ID     string                 `json:"id"`
//	Fields map[string]interface{} `json:"fields"`
//}

type PurchaseFields struct {
	MobileNumber string `json:"Mobile Number,omitempty"`
	Status       string `json:"Status,omitempty"`

	//Other   map[string]interface{}  `json:"-"`
}

type Record struct {
	ID        string         `json:"id"`
	Fields    PurchaseFields `json:"fields"`
	CreatedAt string         `json:"createdTime"`
}

type ListResponse struct {
	Records []Record `json:"records"`
}
