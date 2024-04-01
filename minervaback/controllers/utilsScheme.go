package controllers

type DefaultResponse struct {
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
