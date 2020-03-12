package models

import "math/big"

// Currency model.
type Currency struct {
	ID                uint16     `json:"id"`
	Currency          string     `json:"currency"`
	FullName          string     `json:"fullname"`
	LogoURL           string     `json:"logo_url"`
	Chain             string     `json:"chain"`
	Contract          string     `json:"contract"`
	Decimals          uint8      `json:"decimals"`
	Confirm           uint8      `json:"confirm"`
	DepositEnabled    bool       `json:"deposit_enabled"`
	WithdrawEnabled   bool       `json:"withdraw_enabled"`
	DepositMinAmount  *big.Float `json:"deposit_min_amount"`
	WithdrawMinAmount *big.Float `json:"withdraw_min_amount"`
	WithdrawMaxAmount *big.Float `json:"withdraw_max_amount"`
	WithdrawFeeRate   *big.Float `json:"withdraw_fee_rate"`

	WithdrawPrecision uint8 `json:"withdraw_precision"`
}
