package model

type StandarResponse struct {
	Status  int         `json:"status"`
	Error   string      `json:"error"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}