package forpay

import (
	"strconv"

	"github.com/hactrox/forpay-sdk-go/forpay/request"
	"github.com/hactrox/forpay-sdk-go/forpay/response"
	"github.com/hactrox/forpay-sdk-go/models"
)

// LockBalanceRequest .
type LockBalanceRequest struct {
	*request.BaseRequest
}

// LockBalanceResponse .
type LockBalanceResponse struct {
	*response.BaseResponse
	Data *models.LockBalance `json:"data"`
}

// UnlockBalanceRequest .
type UnlockBalanceRequest struct {
	*request.BaseRequest
}

// UnlockBalanceResponse .
type UnlockBalanceResponse struct {
	*response.BaseResponse
	Data *models.UnlockBalance `json:"data"`
}

// GetLockedBalanceRequest .
type GetLockedBalanceRequest struct {
	*request.BaseRequest
}

// GetLockedBalanceResponse .
type GetLockedBalanceResponse struct {
	*response.BaseResponse
	Data *models.LockedBalance `json:"data"`
}

// GetLockedBalancesRequest .
type GetLockedBalancesRequest struct {
	*request.BaseRequest
}

// GetLockedBalancesResponse .
type GetLockedBalancesResponse struct {
	*response.BaseResponse
	Data []models.LockedBalance `json:"data"`
}

// CreateLockBalanceRequest creates lock balance request.
func CreateLockBalanceRequest(clientToken string, currencyID uint16, walletID uint64, amount string) *LockBalanceRequest {
	requestFields := map[string]interface{}{
		"client_token": clientToken,
		"currency_id":  currencyID,
		"wallet_id":    walletID,
		"amount":       amount,
	}

	req := request.Post("balance/lock", requestFields)
	return &LockBalanceRequest{BaseRequest: req}
}

// LockBalance invokes the 'POST /{version}/balance/lock' API.
func (client *Client) LockBalance(req *LockBalanceRequest) (*LockBalanceResponse, error) {
	resp := &LockBalanceResponse{BaseResponse: &response.BaseResponse{}}
	err := client.RequestWithAuth(req, resp)
	return resp, err
}

// CreateUnlockBalanceRequest creates lock balance request.
func CreateUnlockBalanceRequest(clientToken string, currencyID uint16, walletID uint64, amount string) *UnlockBalanceRequest {
	requestFields := map[string]interface{}{
		"client_token": clientToken,
		"currency_id":  currencyID,
		"wallet_id":    walletID,
		"amount":       amount,
	}

	req := request.Post("balance/unlock", requestFields)
	return &UnlockBalanceRequest{BaseRequest: req}
}

// UnlockBalance invokes the 'POST /{version}/balance/unlock' API.
func (client *Client) UnlockBalance(req *UnlockBalanceRequest) (*UnlockBalanceResponse, error) {
	resp := &UnlockBalanceResponse{BaseResponse: &response.BaseResponse{}}
	err := client.RequestWithAuth(req, resp)
	return resp, err
}

// CreateGetLockedBalanceRequest creates lock balance request.
func CreateGetLockedBalanceRequest(currencyID uint16, walletID uint64) *GetLockedBalanceRequest {
	req := request.Get("balance/locked")
	req.AddQueryParam("currency_id", strconv.Itoa(int(currencyID)))
	req.AddQueryParam("wallet_id", strconv.FormatUint(walletID, 10))

	return &GetLockedBalanceRequest{BaseRequest: req}
}

// GetLockedBalance invokes the 'POST /{version}/balance/locked?wallet_id={walletID}&currency_id={currencyID}' API.
func (client *Client) GetLockedBalance(req *GetLockedBalanceRequest) (*GetLockedBalanceResponse, error) {
	resp := &GetLockedBalanceResponse{BaseResponse: &response.BaseResponse{}}
	err := client.RequestWithAuth(req, resp)
	return resp, err
}

// CreateGetLockedBalancesRequest creates lock balance request.
func CreateGetLockedBalancesRequest(walletID uint64) *GetLockedBalanceRequest {
	req := request.Get("balance/locked")
	req.AddQueryParam("wallet_id", strconv.FormatUint(walletID, 10))

	return &GetLockedBalanceRequest{BaseRequest: req}
}

// GetLockedBalances invokes the 'POST /{version}/balance/locked' API.
func (client *Client) GetLockedBalances(req *GetLockedBalanceRequest) (*GetLockedBalancesResponse, error) {
	resp := &GetLockedBalancesResponse{BaseResponse: &response.BaseResponse{}}
	err := client.RequestWithAuth(req, resp)
	return resp, err
}
