package forpay

import (
	"strconv"

	"github.com/forpay/forpay-sdk-go/forpay/request"
	"github.com/forpay/forpay-sdk-go/forpay/response"
	"github.com/forpay/forpay-sdk-go/models"
)

// GetCurrencyRequest .
type GetCurrencyRequest struct {
	*request.BaseRequest
}

// GetCurrencyResponse .
type GetCurrencyResponse struct {
	*response.BaseResponse
	Data *models.Currency `json:"data"`
}

// GetCurrenciesRequest .
type GetCurrenciesRequest struct {
	*request.BaseRequest
}

// GetCurrenciesResponse .
type GetCurrenciesResponse struct {
	*response.BaseResponse
	Data []models.Currency `json:"data"`
}

// CreateGetCurrencyRequest creates get currency request.
func CreateGetCurrencyRequest(currencyID uint16) *GetCurrencyRequest {
	req := request.Get("currencies")
	req.AddQueryParam("currency_id", strconv.Itoa(int(currencyID)))

	return &GetCurrencyRequest{BaseRequest: req}
}

// GetCurrency invokes the 'GET /{version}/currencies?currency_id={currencyID}' API.
func (client *Client) GetCurrency(req *GetCurrencyRequest) (*GetCurrencyResponse, error) {
	resp := &GetCurrencyResponse{BaseResponse: &response.BaseResponse{}}
	err := client.RequestWithAuth(req, resp)
	return resp, err
}

// CreateGetCurrenciesRequest creates get currencies requesy.
func CreateGetCurrenciesRequest() *GetCurrenciesRequest {
	req := request.Get("currencies")
	return &GetCurrenciesRequest{BaseRequest: req}
}

// GetCurrencies invokes the 'GET /{version}/currencies' API.
func (client *Client) GetCurrencies(req *GetCurrenciesRequest) (*GetCurrenciesResponse, error) {
	resp := &GetCurrenciesResponse{BaseResponse: &response.BaseResponse{}}
	err := client.RequestWithAuth(req, resp)

	return resp, err
}
