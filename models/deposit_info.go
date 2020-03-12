package models

// DepositInfo model.
type DepositInfo struct {
	Currency *Currency `json:"currency"`
	Address  string    `json:"address"`
}
