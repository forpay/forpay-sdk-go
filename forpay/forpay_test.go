package forpay

import (
	"os"
	"testing"

	"github.com/hactrox/forpay-sdk-go/forpay/request"
	"github.com/hactrox/forpay-sdk-go/forpay/response"
)

func setup(t *testing.T) *Client {
	if testing.Short() {
		t.SkipNow()
	}

	appID := os.Getenv("TEST_APPID")
	keyID := os.Getenv("TEST_KEYID")
	filePath := os.Getenv("TEST_PRIVATE_KEY_FILE")

	client, err := NewClientWithRSA(appID, keyID, filePath)
	if err != nil {
		t.Fatal(err)
	}

	return client
}

func testMode(req request.ForpayRequest) {
	req.SetScheme(os.Getenv("TEST_SCHEME"))
	req.SetDomain(os.Getenv("TEST_DOMAIN"))
	req.SetPort(os.Getenv("TEST_PORT"))
}

func checkErr(t *testing.T, err error) {
	if err == nil {
		return
	}

	if errResp, ok := err.(*response.Error); ok {
		if !errResp.IsBusinessFailed() {
			t.Fatal(errResp)
		}

		debug(errResp.Error())
		t.SkipNow()
		return
	}

	t.Fatal(err)
}
