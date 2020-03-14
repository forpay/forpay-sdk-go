package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"
)

// HTTP Constants.
const (
	HEAD = "HEAD"
	GET  = "GET"
	POST = "POST"

	JSON = "application/json"

	HTTP  = "HTTP"
	HTTPS = "HTTPS"
)

var _ ForpayRequest = &BaseRequest{}

// ForpayRequest defines forpay request structure.
type ForpayRequest interface {
	GetScheme() string
	GetDomain() string
	GetPort() string
	GetMethod() string
	GetHeaders() map[string]string
	GetQueryParams() map[string]string
	GetBody() []byte
	GetBodyReader() io.Reader
	GetAPIVersion() string
	GetEndpoint() string
	GetReadTimeout() time.Duration
	GetConnectTimeout() time.Duration

	GetTimestamp() int64
	GetNonce() string

	SetScheme(scheme string)
	SetDomain(domain string)
	SetPort(port string)
	SetMethod(method string)
	SetAPIVersion(version string)

	BuildQueries() string
	BuildURL() string

	SetTimestamp(timestamp int64)
	SetNonce(nonce string)

	AddHeaderParam(key, value string)
	AddQueryParam(key, value string)

	setRequestBody(body []byte)
	setRequestBodyFields(fields map[string]interface{})
}

// BaseRequest implements ForpayRequest interface.
type BaseRequest struct {
	Scheme string
	Method string
	Domain string
	Port   string

	ReadTimeout    time.Duration
	ConnectTimeout time.Duration

	version  string
	endpoint string

	timestamp int64
	nonce     string

	QueryParams map[string]string
	Headers     map[string]string
	RequestBody []byte
}

// Get creates HTTP GET request.
func Get(endpoint string) *BaseRequest {
	return initBaseRequest(GET, endpoint)
}

// Post creates HTTP POST request.
func Post(endpoint string, requestFields map[string]interface{}) *BaseRequest {
	baseRequest := initBaseRequest(POST, endpoint)
	baseRequest.setRequestBodyFields(requestFields)
	baseRequest.AddHeaderParam("Content-Type", JSON)

	return baseRequest
}

func initBaseRequest(method, endpoint string) *BaseRequest {
	return &BaseRequest{
		Scheme:      HTTPS,
		Method:      method,
		Domain:      "api.forpay.pro",
		QueryParams: make(map[string]string),
		Headers:     make(map[string]string),

		version:  "v1",
		endpoint: endpoint,
	}
}

// GetScheme returns request scheme.
func (baseRequest *BaseRequest) GetScheme() string {
	return baseRequest.Scheme
}

// GetDomain returns request domain.
func (baseRequest *BaseRequest) GetDomain() string {
	return baseRequest.Domain
}

// GetPort returns request port.
func (baseRequest *BaseRequest) GetPort() string {
	return baseRequest.Port
}

// GetMethod returns request method.
func (baseRequest *BaseRequest) GetMethod() string {
	return baseRequest.Method
}

// GetHeaders returns request headers.
func (baseRequest *BaseRequest) GetHeaders() map[string]string {
	return baseRequest.Headers
}

// GetQueryParams returns request query params.
func (baseRequest *BaseRequest) GetQueryParams() map[string]string {
	return baseRequest.QueryParams
}

// GetBody returns request body bytes.
func (baseRequest *BaseRequest) GetBody() []byte {
	return baseRequest.RequestBody
}

// GetBodyReader returns request body as io.Reader.
func (baseRequest *BaseRequest) GetBodyReader() io.Reader {
	return bytes.NewBuffer(baseRequest.RequestBody)
}

// GetAPIVersion returns request api version.
func (baseRequest *BaseRequest) GetAPIVersion() string {
	return baseRequest.version
}

// GetEndpoint returns request api endpoint.
func (baseRequest *BaseRequest) GetEndpoint() string {
	return baseRequest.endpoint
}

// GetReadTimeout returns request read timeout.
func (baseRequest *BaseRequest) GetReadTimeout() time.Duration {
	return baseRequest.ReadTimeout
}

// GetConnectTimeout returns request connect timeout.
func (baseRequest *BaseRequest) GetConnectTimeout() time.Duration {
	return baseRequest.ConnectTimeout
}

// GetTimestamp returns request timestamp.
func (baseRequest *BaseRequest) GetTimestamp() int64 {
	return baseRequest.timestamp
}

// GetNonce returns request nonce.
func (baseRequest *BaseRequest) GetNonce() string {
	return baseRequest.nonce
}

// SetScheme sets request scheme.
func (baseRequest *BaseRequest) SetScheme(scheme string) {
	baseRequest.Scheme = scheme
}

// SetDomain sets request domain.
func (baseRequest *BaseRequest) SetDomain(domain string) {
	baseRequest.Domain = domain
}

// SetPort sets request port.
func (baseRequest *BaseRequest) SetPort(port string) {
	baseRequest.Port = port
}

// SetMethod sets request method.
func (baseRequest *BaseRequest) SetMethod(method string) {
	baseRequest.Method = method
}

// SetAPIVersion sets request api version.
func (baseRequest *BaseRequest) SetAPIVersion(version string) {
	baseRequest.version = version
}

// BuildQueries returns encoded request queries.
func (baseRequest *BaseRequest) BuildQueries() string {
	queryParams := baseRequest.GetQueryParams()
	if len(queryParams) == 0 {
		return ""
	}

	urlEncoder := url.Values{}

	for key, value := range queryParams {
		urlEncoder.Add(key, value)
	}

	return "?" + urlEncoder.Encode()
}

// BuildURL returns full request URL.
func (baseRequest *BaseRequest) BuildURL() string {
	url := fmt.Sprintf("%s://%s", strings.ToLower(baseRequest.Scheme), baseRequest.Domain)
	if baseRequest.Port != "" {
		url = fmt.Sprintf("%s:%s", url, baseRequest.Port)
	}

	url = fmt.Sprintf("%s/%s/%s", url, baseRequest.version, baseRequest.endpoint)

	return url + baseRequest.BuildQueries()
}

// SetTimestamp sets request timestamp.
func (baseRequest *BaseRequest) SetTimestamp(timestamp int64) {
	baseRequest.timestamp = timestamp
}

// SetNonce sets request nonce.
func (baseRequest *BaseRequest) SetNonce(nonce string) {
	baseRequest.nonce = nonce
}

// AddHeaderParam adds request header to the request.
func (baseRequest *BaseRequest) AddHeaderParam(key, value string) {
	baseRequest.Headers[key] = value
}

// AddQueryParam adds query param to the request.
func (baseRequest *BaseRequest) AddQueryParam(key, value string) {
	baseRequest.QueryParams[key] = value
}

func (baseRequest *BaseRequest) setRequestBody(body []byte) {
	baseRequest.RequestBody = body
}

func (baseRequest *BaseRequest) setRequestBodyFields(fields map[string]interface{}) {
	requestBody, _ := json.Marshal(fields)
	baseRequest.RequestBody = requestBody
}
