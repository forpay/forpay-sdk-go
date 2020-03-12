package models

// Transfer model.
type Transfer struct {
	TransactionID       uint64 `json:"transaction_id"`
	IsIdempotentRequest bool   `json:"is_idempotent_request"`
}
