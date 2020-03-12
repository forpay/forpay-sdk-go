package forpay

import (
	"testing"
)

func TestGetCurrency(t *testing.T) {
	client := setup(t)

	req := CreateGetCurrencyRequest(1)
	testMode(req)

	resp, err := client.GetCurrency(req)
	checkErr(t, err)

	if resp.Data == nil {
		t.Fatal("failed to get response data")
	}
	if resp.Data.ID == 0 ||
		resp.Data.Chain == "" ||
		resp.Data.Currency == "" {
		t.Fatal("failed to get currency info")
	}
}

func TestGetCurrencies(t *testing.T) {
	client := setup(t)

	req := CreateGetCurrenciesRequest()
	testMode(req)

	_, err := client.GetCurrencies(req)
	checkErr(t, err)
}
