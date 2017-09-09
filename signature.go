package amazon_alexa_web_services_api_client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
)

// Method used to sign
const (
	SignatureMethod  = "HmacSHA256"
	SignatureVersion = "2"
)

func sign(host, secretKey, query string) string {
	signature := fmt.Sprintf("GET\n%s\n/\n%s", strings.ToLower(host), query)

	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(signature))

	return url.QueryEscape(base64.StdEncoding.EncodeToString(h.Sum(nil)))
}
