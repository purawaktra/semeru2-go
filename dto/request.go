package dto

type Request struct {
	RequestId string `json:"request_id"`
	Data      any    `json:"data"`
}

type RequestCredential struct {
	AccountId    int    `json:"account_id"`
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}
