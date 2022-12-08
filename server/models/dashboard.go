package models

type DashboardSum struct {
	Customers      int       `json:"customers"`
	Contracts      int       `json:"contracts"`
	ContractAmount float64   `json:"contractAmount"`
	Products       int       `json:"products"`
	Date           []string  `json:"date"`
	Amount         []float64 `json:"amount"`
}
