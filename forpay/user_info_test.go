package forpay

import (
	"testing"
)

func TestSyncUserInfo(t *testing.T) {
	client := setup(t)

	userID := "123"
	req := CreateSyncUserInfoRequest(userID)
	testMode(req)

	resp, err := client.SyncUserInfo(req)
	checkErr(t, err)

	if resp.Data == nil {
		t.Fatal("failed to get response data")
	}

	if resp.Data.UserID != userID {
		t.Fatalf("incorrect userID from 'SyncUserInfo', get %s, want %s", resp.Data.UserID, userID)
	}
	if resp.Data.WalletID == 0 {
		t.Fatal("failed toget walletID")
	}
}
