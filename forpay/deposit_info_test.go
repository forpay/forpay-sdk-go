package forpay

import (
	"os"
	"strconv"
	"testing"
)

func TestGetDepositInfo(t *testing.T) {
	client := setup(t)

	walletID, _ := strconv.ParseUint(os.Getenv("TEST_WALLETID"), 10, 64)
	currencyID := uint16(1)
	req := CreateGetDepositInfoRequest(walletID, currencyID)
	testMode(req)

	resp, err := client.GetDepositInfo(req)
	checkErr(t, err)

	if resp.Data == nil {
		t.Fatal("failed to get response data")
	}

	if resp.Data.Address == "" {
		t.Fatal("failed to get address")
	}
	if resp.Data.Currency == nil {
		t.Fatal("failed to get currency info")
	}
}
