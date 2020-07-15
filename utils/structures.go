package utils

type Response struct {
	Message interface{} `json:"message"`
	Error   interface{} `json:"error"`
}
