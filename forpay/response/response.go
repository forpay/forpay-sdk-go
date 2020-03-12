package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var _ ForpayResponse = &BaseResponse{}

// ForpayResponse defines forpay response structure.
type ForpayResponse interface {
	IsSuccess() bool

	GetError() error
	GetHTTPStatus() int
	GetHTTPHeaders() map[string][]string
	GetHTTPContentBytes() []byte
	GetOriginHTTPResponse() *http.Response

	loadResponse(httpResponse *http.Response) error
}

// BaseResponse implements ForpayResponse interface.
type BaseResponse struct {
	ErrorResponse

	httpStatus         int
	httpHeaders        map[string][]string
	httpContentBytes   []byte
	httpContentString  string
	originHTTPResponse *http.Response
}

// IsSuccess tells if the http request succeed.
func (baseResponse *BaseResponse) IsSuccess() bool {
	httpStatus := baseResponse.GetHTTPStatus()
	return httpStatus >= 200 && httpStatus < 300
}

// GetError returns error if request failed.
func (baseResponse *BaseResponse) GetError() error {
	errResp := baseResponse.ErrorResponse
	if errResp.Code == 0 && errResp.Msg == "Success" {
		return nil
	}

	return errResp
}

// GetHTTPStatus returns http status of the response.
func (baseResponse *BaseResponse) GetHTTPStatus() int {
	return baseResponse.httpStatus
}

// GetHTTPHeaders returns http headers from response.
func (baseResponse *BaseResponse) GetHTTPHeaders() map[string][]string {
	return baseResponse.httpHeaders
}

// GetHTTPContentBytes returns http response body bytes.
func (baseResponse *BaseResponse) GetHTTPContentBytes() []byte {
	return baseResponse.httpContentBytes
}

// GetOriginHTTPResponse returns raw http response.
func (baseResponse *BaseResponse) GetOriginHTTPResponse() *http.Response {
	return baseResponse.originHTTPResponse
}

func (baseResponse *BaseResponse) loadResponse(httpResponse *http.Response) error {
	defer httpResponse.Body.Close()
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}

	baseResponse.httpStatus = httpResponse.StatusCode
	baseResponse.httpHeaders = httpResponse.Header
	baseResponse.httpContentBytes = body
	baseResponse.httpContentString = string(body)
	baseResponse.originHTTPResponse = httpResponse

	return nil
}

// Unmarshal parses the JSON-encodeded data and stores the value in ForpayResponse.
func Unmarshal(resp ForpayResponse, httpResponse *http.Response) error {
	err := resp.loadResponse(httpResponse)
	if err != nil {
		return err
	}

	respBody := resp.GetHTTPContentBytes()

	if !resp.IsSuccess() {
		errMsg := "RequestError"
		errMsg += fmt.Sprintf("\nHttpStatus: %d", resp.GetHTTPStatus())

		if respBody != nil {
			errMsg += fmt.Sprintf("\nResponse: %s", string(respBody))
		}

		return errors.New(errMsg)
	}

	if len(respBody) == 0 {
		return nil
	}

	err = json.Unmarshal(resp.GetHTTPContentBytes(), &resp)
	if err != nil {
		return err
	}

	if err := resp.GetError(); err != nil {
		return err
	}

	return json.Unmarshal(respBody, resp)
}
