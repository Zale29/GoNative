package models

type Response struct {
	Status  int         `json:"status"`
	Message int         `json:"message"`
	Data    interface{} `json:"data"`
}
