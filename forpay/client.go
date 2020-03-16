package forpay

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/forpay/forpay-sdk-go/forpay/request"
	"github.com/forpay/forpay-sdk-go/forpay/response"
	"github.com/forpay/forpay-sdk-go/utils"
)

var debug utils.Debug

func init() {
	debug = utils.InitDebug()
}

// Version is the SDK version of the package.
const Version = "0.9.1"

var (
	defaultConnectTimeout = 5 * time.Second
	defaultReadTimeout    = 10 * time.Second

	defaultUserAgent = fmt.Sprintf("Forpay (%s; %s) Golang/%s SDK/%s", runtime.GOOS, runtime.GOARCH, strings.Trim(runtime.Version(), "go"), Version)
)

// Client is the forpay SDK client.
type Client struct {
	appID string
	keyID string

	readTimeout    time.Duration
	connectTimeout time.Duration

	config *Config

	httpClient *http.Client

	privKey *rsa.PrivateKey
}

// NewClientWithRSA creates new forpay SDK client with rsa private keyfile loaded.
func NewClientWithRSA(appID, keyID, filePath string) (*Client, error) {
	client := &Client{
		appID: appID,
		keyID: keyID,
	}

	config := DefaultConfig()
	client.LoadConfig(config)

	err := client.loadPrivateKey(filePath)
	return client, err
}

func (client *Client) loadPrivateKey(filePath string) error {
	keyData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return errors.New("invalid private key file")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}

	client.privKey = priv

	return nil
}

// LoadConfig applies configs for client.
func (client *Client) LoadConfig(config *Config) {
	client.config = config
	client.httpClient = &http.Client{}

	if config.Timeout > 0 {
		client.httpClient.Timeout = config.Timeout
	}
}

// SetReadTimeout sets read timeout for clinet.
func (client *Client) SetReadTimeout(readTimeout time.Duration) {
	client.readTimeout = readTimeout
}

// SetConnectTimeout sets connect timeout for client.
func (client *Client) SetConnectTimeout(connectTimeout time.Duration) {
	client.connectTimeout = connectTimeout
}

// GetReadTimeout returns read timeout.
func (client *Client) GetReadTimeout() time.Duration {
	return client.readTimeout
}

// GetConnectTimeout returns connect timeout.
func (client *Client) GetConnectTimeout() time.Duration {
	return client.connectTimeout
}

func (client *Client) getTimeout(req request.ForpayRequest) (time.Duration, time.Duration) {
	readTimeout := defaultReadTimeout
	connectTimeout := defaultConnectTimeout

	reqReadTimeout := req.GetReadTimeout()
	reqConnectTimeout := req.GetConnectTimeout()

	if reqReadTimeout != 0*time.Millisecond {
		readTimeout = reqReadTimeout
	} else if client.readTimeout != 0*time.Millisecond {
		readTimeout = client.readTimeout
	} else if client.httpClient.Timeout != 0 {
		readTimeout = client.httpClient.Timeout
	}

	if reqConnectTimeout != 0*time.Millisecond {
		connectTimeout = reqConnectTimeout
	} else if client.connectTimeout != 0*time.Millisecond {
		connectTimeout = client.connectTimeout
	}
	return readTimeout, connectTimeout
}

// Timeout sets DialContext func with timeout.
func Timeout(connectTimeout time.Duration) func(cxt context.Context, net, addr string) (c net.Conn, err error) {
	return func(ctx context.Context, network, address string) (net.Conn, error) {
		return (&net.Dialer{
			Timeout:   connectTimeout,
			DualStack: true,
		}).DialContext(ctx, network, address)
	}
}

func (client *Client) setTimeout(req request.ForpayRequest) {
	readTimeout, connectTimeout := client.getTimeout(req)
	client.httpClient.Timeout = readTimeout
	if trans, ok := client.httpClient.Transport.(*http.Transport); ok && trans != nil {
		trans.DialContext = Timeout(connectTimeout)
		client.httpClient.Transport = trans
	} else if client.httpClient.Transport == nil {
		client.httpClient.Transport = &http.Transport{
			DialContext: Timeout(connectTimeout),
		}
	}
}

// Request sends api request without authorization header.
func (client *Client) Request(req request.ForpayRequest, resp response.ForpayResponse) error {
	httpRequest, err := client.buildRequest(req)
	if err != nil {
		return err
	}

	for key, value := range httpRequest.Header {
		debug("%s: %s", key, value)
	}

	return client.request(req, resp, httpRequest)
}

// RequestWithAuth sends api request with authorization header.
func (client *Client) RequestWithAuth(req request.ForpayRequest, resp response.ForpayResponse) error {
	httpRequest, err := client.buildRequestWithAuth(req)
	if err != nil {
		return err
	}

	for key, value := range httpRequest.Header {
		debug("%s: %s", key, value)
	}

	return client.request(req, resp, httpRequest)
}

func (client *Client) request(req request.ForpayRequest, resp response.ForpayResponse, httpRequest *http.Request) error {
	client.setTimeout(req)
	httpResponse, err := client.httpClient.Do(httpRequest)
	if err != nil {
		return err
	}

	return response.Unmarshal(resp, httpResponse)
}

func (client *Client) buildRequest(req request.ForpayRequest) (*http.Request, error) {
	url := req.BuildURL()
	method := req.GetMethod()
	body := req.GetBody()

	debug("method: %s", method)
	debug("full url: %s", url)
	if body == nil {
		debug("request body: nil")
	} else {
		debug("request body: %s", string(body))
	}

	httpRequest, err := http.NewRequest(method, url, req.GetBodyReader())
	if err != nil {
		return nil, err
	}

	req.AddHeaderParam("X-SDK-Version", Version)
	req.AddHeaderParam("User-Agent", defaultUserAgent)

	for key, value := range req.GetHeaders() {
		httpRequest.Header[key] = []string{value}
	}

	return httpRequest, nil
}

func (client *Client) buildRequestWithAuth(req request.ForpayRequest) (*http.Request, error) {
	if err := signRequest(client.appID, client.keyID, req, client.privKey); err != nil {
		return nil, err
	}

	httpRequest, err := client.buildRequest(req)
	return httpRequest, err
}
