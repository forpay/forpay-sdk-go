package forpay

import (
	"os"
	"strconv"
	"testing"

	"github.com/hactrox/forpay-sdk-go/utils"
)

func TestLockBalance(t *testing.T) {
	client := setup(t)

	clientToken := utils.RandStringRunes(32)
	walletID, _ := strconv.ParseUint(os.Getenv("TEST_WALLETID"), 10, 64)
	currencyID := uint16(1)
	amount := "0.001"

	req := CreateLockBalanceRequest(clientToken, currencyID, walletID, amount)
	testMode(req)

	_, err := client.LockBalance(req)
	checkErr(t, err)
}

func TestUnlockBalance(t *testing.T) {
	client := setup(t)

	clientToken := utils.RandStringRunes(32)
	walletID, _ := strconv.ParseUint(os.Getenv("TEST_WALLETID"), 10, 64)
	currencyID := uint16(1)
	amount := "0.001"

	req := CreateUnlockBalanceRequest(clientToken, currencyID, walletID, amount)
	testMode(req)

	_, err := client.UnlockBalance(req)
	checkErr(t, err)
}

func TestGetLockedBalance(t *testing.T) {
	client := setup(t)

	walletID, _ := strconv.ParseUint(os.Getenv("TEST_WALLETID"), 10, 64)
	currencyID := uint16(1)

	req := CreateGetLockedBalanceRequest(currencyID, walletID)
	testMode(req)

	resp, err := client.GetLockedBalance(req)
	checkErr(t, err)

	if resp.Data == nil {
		t.Fatal("failed to get response data")
	}
	if resp.Data.CurrencyID != currencyID {
		t.Fatalf("incorrect currencyID, get %d, want %d",
			resp.Data.CurrencyID, currencyID)
	}
	if resp.Data.AmountLocked == nil {
		t.Fatal("failed to get amount locked")
	}
}

func TestGetLockedBalances(t *testing.T) {
	client := setup(t)

	walletID, _ := strconv.ParseUint(os.Getenv("TEST_WALLETID"), 10, 64)

	req := CreateGetLockedBalancesRequest(walletID)
	testMode(req)

	resp, err := client.GetLockedBalances(req)
	checkErr(t, err)

	if resp.Data == nil {
		t.Fatal("failed to get response data")
	}
	if len(resp.Data) == 0 {
		t.Fatal("failed to get locked balances")
	}
	if resp.Data[0].AmountLocked == nil ||
		resp.Data[0].CurrencyID == 0 {
		t.Fatal("failed to get locked balance data")
	}
}
