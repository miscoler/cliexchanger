package models

type ExchangeResponse struct {
	Rates map[string]float64 `json:"rates"`
	Error string             `json:"error"`
}
