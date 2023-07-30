package dto

type BodyCredential struct {
	AccountId    int    `json:"account_id"`
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

type Header struct {
	RequestId string `header:"request-id"`
	Host      string `header:"host"`
}
