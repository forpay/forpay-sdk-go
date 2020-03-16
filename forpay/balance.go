package forpay

import (
	"strconv"

	"github.com/forpay/forpay-sdk-go/forpay/request"
	"github.com/forpay/forpay-sdk-go/forpay/response"
	"github.com/forpay/forpay-sdk-go/models"
)

// GetBalanceRequest .
type GetBalanceRequest struct {
	*request.BaseRequest
}

// GetBalanceResponse .
type GetBalanceResponse struct {
	*response.BaseResponse
	Data *models.Balance `json:"data"`
}

// CreateGetBalanceRequest creates get balance request.
func CreateGetBalanceRequest(walletID uint64, currencyID uint16) *GetBalanceRequest {
	req := request.Get("balance")
	req.AddQueryParam("wallet_id", strconv.FormatUint(walletID, 10))
	req.AddQueryParam("currency_id", strconv.Itoa(int(currencyID)))

	return &GetBalanceRequest{BaseRequest: req}
}

// GetBalance invokes the 'GET /{version}/balance' API.
func (client *Client) GetBalance(req *GetBalanceRequest) (*GetBalanceResponse, error) {
	resp := &GetBalanceResponse{BaseResponse: &response.BaseResponse{}}
	err := client.RequestWithAuth(req, resp)
	return resp, err
}
