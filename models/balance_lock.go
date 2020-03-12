package models

import "math/big"

// LockedBalance .
type LockedBalance struct {
	CurrencyID   uint16     `json:"currency_id"`
	AmountLocked *big.Float `json:"amount_locked"`
}
