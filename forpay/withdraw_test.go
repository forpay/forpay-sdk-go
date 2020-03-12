package forpay

import (
	"os"
	"strconv"
	"testing"

	"github.com/hactrox/forpay-sdk-go/utils"
)

func TestWithdraw(t *testing.T) {
	client := setup(t)

	clientToken := utils.RandStringRunes(32)
	walletID, _ := strconv.ParseUint(os.Getenv("TEST_WALLETID"), 10, 64)
	currencyID := uint16(1)
	amount := "0.001"
	address := os.Getenv("TEST_WITHDRAW_ADDR")

	req := CreateWithdrawRequest(clientToken, walletID, currencyID, amount, address)
	testMode(req)

	resp, err := client.Withdraw(req)
	checkErr(t, err)

	if resp.Data == nil {
		t.Fatal("failed to get response data")
	}
	if resp.Data.TransactionID == 0 {
		t.Fatal("failed to get transactionID")
	}
}
