package dto

type BaseResponse struct {
	ResponseId   string `json:"response_id"`
	RequestId    string `json:"request_id"`
	Success      bool   `json:"success"`
	ResponseTime string `json:"response_time"`
}

type ResponseError struct {
	BaseResponse
	Errors any `json:"errors,omitempty"`
}

type ResponseSuccess struct {
	BaseResponse
	Message string `json:"message"`
	Data    any    `json:"data"`
}
