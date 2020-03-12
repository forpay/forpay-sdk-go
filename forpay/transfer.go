package forpay

import (
	"github.com/hactrox/forpay-sdk-go/forpay/request"
	"github.com/hactrox/forpay-sdk-go/forpay/response"
	"github.com/hactrox/forpay-sdk-go/models"
)

// TransferRequest .
type TransferRequest struct {
	*request.BaseRequest
}

// TransferResponse .
type TransferResponse struct {
	*response.BaseResponse
	Data *models.Transfer `json:"data"`
}

// CreateTransferRequest creates transfer request.
func CreateTransferRequest(clientToken string, currencyID uint16, from, to uint64, amount string) *TransferRequest {
	requestFields := map[string]interface{}{
		"client_token": clientToken,
		"currency_id":  currencyID,
		"from":         from,
		"to":           to,
		"amount":       amount,
	}

	req := request.Post("transfer", requestFields)
	return &TransferRequest{BaseRequest: req}
}

// Transfer invokes the 'POST /{version}/transfer' API.
func (client *Client) Transfer(req *TransferRequest) (*TransferResponse, error) {
	resp := &TransferResponse{BaseResponse: &response.BaseResponse{}}
	err := client.RequestWithAuth(req, resp)
	return resp, err
}
