package models

import "math/big"

// Balance model.
type Balance struct {
	CurrencyID uint16     `json:"currency_id"`
	Available  *big.Float `json:"available"`
	Locked     *big.Float `json:"locked"`
	Frozen     *big.Float `json:"frozen"`
}
