package airtable

//type Record struct {
//	ID     string                 `json:"id"`
//	Fields map[string]interface{} `json:"fields"`
//}

type PurchaseFields struct {
	Email    *string `json:"Email,omitempty"`
	FullName *string `json:"Получатель(ФИО),omitempty"`
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
