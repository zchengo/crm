package models

type DashboardSum struct {
	Customers      int        `json:"customers"`
	Contracts      int        `json:"contracts"`
	ContractAmount float64    `json:"contractAmount"`
	Products       int        `json:"products"`
	Date           [7]string  `json:"date"`
	Amount         [7]float64 `json:"amount"`
}
