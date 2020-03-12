package forpay

import (
	"github.com/hactrox/forpay-sdk-go/forpay/request"
	"github.com/hactrox/forpay-sdk-go/forpay/response"
	"github.com/hactrox/forpay-sdk-go/models"
)

// SyncUserInfoRequest .
type SyncUserInfoRequest struct {
	*request.BaseRequest
}

// SyncUserInfoResponse .
type SyncUserInfoResponse struct {
	*response.BaseResponse
	Data *models.UserInfo `json:"data"`
}

// CreateSyncUserInfoRequest creates sync user info request.
func CreateSyncUserInfoRequest(userID string) *SyncUserInfoRequest {
	requestFields := map[string]interface{}{
		"user_id": userID,
	}

	req := request.Post("user_info", requestFields)
	return &SyncUserInfoRequest{BaseRequest: req}
}

// SyncUserInfo invokes the 'POST /{version}/user_info' API.
func (client *Client) SyncUserInfo(req *SyncUserInfoRequest) (*SyncUserInfoResponse, error) {
	resp := &SyncUserInfoResponse{BaseResponse: &response.BaseResponse{}}
	err := client.RequestWithAuth(req, resp)
	return resp, err
}
