package forpay

import (
	"github.com/hactrox/forpay-sdk-go/forpay/request"
	"github.com/hactrox/forpay-sdk-go/forpay/response"
	"github.com/hactrox/forpay-sdk-go/models"
)

// WithdrawRequest .
type WithdrawRequest struct {
	*request.BaseRequest
}

// WithdrawResponse .
type WithdrawResponse struct {
	*response.BaseResponse
	Data *models.Withdraw `json:"data"`
}

// CreateWithdrawRequest creates withdraw request.
func CreateWithdrawRequest(clientToken string, walletID uint64, currencyID uint16, amount, address string) *WithdrawRequest {
	requestFields := map[string]interface{}{
		"client_token": clientToken,
		"wallet_id":    walletID,
		"currency_id":  currencyID,
		"amount":       amount,
		"address":      address,
	}

	req := request.Post("withdraw", requestFields)
	return &WithdrawRequest{BaseRequest: req}
}

// Withdraw invokes the 'POST /{version}/withdraw' API.
func (client *Client) Withdraw(req *WithdrawRequest) (*WithdrawResponse, error) {
	resp := &WithdrawResponse{BaseResponse: &response.BaseResponse{}}
	err := client.RequestWithAuth(req, resp)
	return resp, err
}
