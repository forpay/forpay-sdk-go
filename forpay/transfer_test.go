package forpay

import (
	"os"
	"strconv"
	"testing"

	"github.com/hactrox/forpay-sdk-go/utils"
)

func TestTransfer(t *testing.T) {
	client := setup(t)

	clientToken := utils.RandStringRunes(32)
	from, _ := strconv.ParseUint(os.Getenv("TEST_TRANSFER_FROM"), 10, 64)
	to, _ := strconv.ParseUint(os.Getenv("TEST_TRANSFER_TO"), 10, 64)
	currencyID := uint16(1)
	amount := "0.001"

	req := CreateTransferRequest(clientToken, currencyID, from, to, amount)
	testMode(req)

	resp, err := client.Transfer(req)
	checkErr(t, err)

	if resp.Data == nil {
		t.Fatal("failed to get response data")
	}
	if resp.Data.TransactionID == 0 {
		t.Fatal("failed to get transactionID")
	}
	if resp.Data.IsIdempotentRequest {
		t.Fatal("get is_idempotent_request=false, want true")
	}
}
