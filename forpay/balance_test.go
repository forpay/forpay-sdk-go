package forpay

import (
	"os"
	"strconv"
	"testing"
)

func TestGetBalance(t *testing.T) {
	client := setup(t)

	walletID, _ := strconv.ParseUint(os.Getenv("TEST_WALLETID"), 10, 64)
	currencyID := uint16(1)
	req := CreateGetBalanceRequest(walletID, currencyID)
	testMode(req)

	resp, err := client.GetBalance(req)
	checkErr(t, err)

	balance := resp.Data

	if balance == nil {
		t.Fatal("failed to get response data")
	}

	if balance.CurrencyID == 0 {
		t.Fatal("failed to get currencyID")
	}
	if balance.Available == nil ||
		balance.Locked == nil ||
		balance.Frozen == nil {
		t.Fatal("failed to get balance info")
	}
}
