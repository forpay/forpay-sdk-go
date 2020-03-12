package forpay

import (
	"strconv"

	"github.com/hactrox/forpay-sdk-go/forpay/request"
	"github.com/hactrox/forpay-sdk-go/forpay/response"
	"github.com/hactrox/forpay-sdk-go/models"
)

// GetDepositInfoRequest .
type GetDepositInfoRequest struct {
	*request.BaseRequest
}

// GetDepositInfoResponse .
type GetDepositInfoResponse struct {
	*response.BaseResponse
	Data *models.DepositInfo `json:"data"`
}

// CreateGetDepositInfoRequest creates get deposit info request.
func CreateGetDepositInfoRequest(walletID uint64, currencyID uint16) *GetDepositInfoRequest {
	req := request.Get("deposit_info")
	req.AddQueryParam("wallet_id", strconv.FormatUint(walletID, 10))
	req.AddQueryParam("currency_id", strconv.Itoa(int(currencyID)))

	return &GetDepositInfoRequest{BaseRequest: req}
}

// GetDepositInfo invokes the 'GET /{version}/deposit_info' API.
func (client *Client) GetDepositInfo(req *GetDepositInfoRequest) (*GetDepositInfoResponse, error) {
	resp := &GetDepositInfoResponse{BaseResponse: &response.BaseResponse{}}
	err := client.RequestWithAuth(req, resp)
	return resp, err
}
