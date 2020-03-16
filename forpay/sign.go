package forpay

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/forpay/forpay-sdk-go/forpay/request"
	"github.com/forpay/forpay-sdk-go/utils"
)

var md5Sum = md5.New()

func signRequest(appID, keyID string, req request.ForpayRequest, privKey *rsa.PrivateKey) error {
	setAuthHeaders(appID, req)
	return sign(appID, keyID, req, privKey)
}

func setAuthHeaders(appID string, req request.ForpayRequest) {
	req.SetTimestamp(time.Now().UnixNano() / 1e6)
	req.SetNonce(utils.RandStringRunes(32))

	req.AddHeaderParam("X-Request-AppId", appID)
	req.AddHeaderParam("X-Request-Timestamp", strconv.FormatInt(req.GetTimestamp(), 10))
	req.AddHeaderParam("X-Request-Nonce", req.GetNonce())
}

func sign(appID, keyID string, req request.ForpayRequest, privKey *rsa.PrivateKey) error {
	stringToSign := prepareStringToSign(appID, req)
	debug(strings.ReplaceAll(stringToSign, "\n", "\\n"))

	hash := crypto.SHA256
	rng := rand.Reader
	h := sha256.New()
	h.Write([]byte(stringToSign))
	hashed := h.Sum(nil)

	cipher, err := rsa.SignPKCS1v15(rng, privKey, hash, hashed)
	if err != nil {
		return err
	}

	signature := base64.StdEncoding.EncodeToString(cipher)
	auth := fmt.Sprintf("SHA256-RSA %s:%s", keyID, signature)

	req.AddHeaderParam("Authorization", auth)

	return nil
}

func prepareStringToSign(appID string, req request.ForpayRequest) string {
	content := getContent(req)
	return fmt.Sprintf("%s\n%d\n%s\n%s\n%s",
		req.GetMethod(),
		req.GetTimestamp(),
		appID,
		req.GetNonce(),
		calContentMD5(content),
	)
}

func getContent(req request.ForpayRequest) string {
	switch req.GetMethod() {
	case request.HEAD:
		return ""
	case request.GET:
		return getContentFromQuery(req.GetQueryParams())
	case request.POST:
		return getContentFromBody(req.GetBody())
	default:
		panic("not supported yet")
	}
}

func getContentFromQuery(queryParams map[string]string) string {
	var keys []string
	for key := range queryParams {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var params []string
	for _, key := range keys {
		value := queryParams[key]
		p := fmt.Sprintf("%s=%s",
			popStandardURLEncode(key),
			popStandardURLEncode(value))
		params = append(params, p)
	}

	content := strings.Join(params, "&")
	return content
}

func popStandardURLEncode(str string) string {
	result := strings.Replace(str, "+", "%20", -1)
	result = strings.Replace(result, "*", "%2A", -1)
	result = strings.Replace(result, "%7E", "~", -1)
	return result
}

func getContentFromBody(body []byte) string {
	return string(body)
}

func calContentMD5(content string) string {
	md5Sum.Reset()
	md5Sum.Write([]byte(content))
	bytes := md5Sum.Sum(nil)

	return base64.StdEncoding.EncodeToString(bytes)
}
