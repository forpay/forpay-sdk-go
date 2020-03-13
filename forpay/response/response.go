package response

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var _ ForpayResponse = &BaseResponse{}

// ForpayResponse defines forpay response structure.
type ForpayResponse interface {
	IsSuccess() bool

	GetHTTPStatus() int
	GetHTTPHeaders() map[string][]string
	GetHTTPContentBytes() []byte
	GetOriginHTTPResponse() *http.Response

	getError() error

	loadResponse(httpResponse *http.Response) error
}

// BaseResponse implements ForpayResponse interface.
type BaseResponse struct {
	httpStatus         int
	httpHeaders        map[string][]string
	httpContentBytes   []byte
	httpContentString  string
	originHTTPResponse *http.Response
}

func (baseResponse *BaseResponse) String() string {
	resultBuilder := bytes.Buffer{}

	origin := baseResponse.originHTTPResponse
	resultBuilder.WriteString(fmt.Sprintf("%s %s\n", origin.Proto, origin.Status))

	for key, values := range baseResponse.httpHeaders {
		resultBuilder.WriteString(key + ": " + strings.Join(values, ";") + "\n")
	}

	resultBuilder.WriteString("\n" + baseResponse.httpContentString)

	return resultBuilder.String()
}

// IsSuccess tells if the http request succeed.
func (baseResponse *BaseResponse) IsSuccess() bool {
	httpStatus := baseResponse.GetHTTPStatus()
	return httpStatus >= 200 && httpStatus < 300
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

func (baseResponse *BaseResponse) getError() error {
	fields := make(map[string]interface{})

	err := &Error{
		HTTPStatus: baseResponse.GetHTTPStatus(),
	}

	e := json.Unmarshal(baseResponse.GetHTTPContentBytes(), &fields)
	if e != nil {
		return err
	}

	err.Code = int(fields["code"].(float64))
	err.Msg = fields["msg"].(string)

	if err.Code == 0 && err.Msg == "Success" {
		return nil
	}

	err.SubCode = fields["sub_code"].(string)
	err.SubMsg = fields["sub_msg"].(string)

	return err
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

func getRequestError(resp ForpayResponse) error {
	errMsg := "RequestError"
	errMsg += fmt.Sprintf("\nHttpStatus: %d", resp.GetHTTPStatus())

	respBody := resp.GetHTTPContentBytes()
	if respBody != nil {
		errMsg += fmt.Sprintf("\nResponse: %s", string(respBody))
	}

	return errors.New(errMsg)
}

// Unmarshal parses the JSON-encodeded data and stores the value in ForpayResponse.
func Unmarshal(resp ForpayResponse, httpResponse *http.Response) error {
	err := resp.loadResponse(httpResponse)
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return getRequestError(resp)
	}

	respBody := resp.GetHTTPContentBytes()
	if len(respBody) == 0 {
		return nil
	}

	err = json.Unmarshal(resp.GetHTTPContentBytes(), &resp)
	if err != nil {
		return err
	}

	if err := resp.getError(); err != nil {
		return err
	}

	return json.Unmarshal(respBody, resp)
}
