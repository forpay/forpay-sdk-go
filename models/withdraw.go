package models

// Withdraw model.
type Withdraw struct {
	TransactionID       uint64 `json:"transaction_id"`
	IsIdempotentRequest bool   `json:"is_idempotent_request"`
}
