package models

type Response struct {
	Message string      `json:"Message" example:"Equipo 2"`
	Success bool        `json:"Success" example:"false"`
	Data    interface{} `json:"Data"`
}
