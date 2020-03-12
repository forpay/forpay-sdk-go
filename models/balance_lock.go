package models

import "math/big"

// LockBalance model.
type LockBalance struct {
	IsIdempotentRequest bool `json:"is_idempotent_request"`
}

// UnlockBalance model.
type UnlockBalance struct {
	IsIdempotentRequest bool `json:"is_idempotent_request"`
}

// LockedBalance model.
type LockedBalance struct {
	CurrencyID   uint16     `json:"currency_id"`
	AmountLocked *big.Float `json:"amount_locked"`
}
