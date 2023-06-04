package models

type Response struct {
	Message string      `json:"Message"`
	Success bool        `json:"Success"`
	Data    interface{} `json:"Data"`
}
