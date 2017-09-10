package awis

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type RequestType uint8

// API Request types
// See documentation at http://docs.aws.amazon.com/AlexaWebInfoService/latest/
const (
	UrlInfo RequestType = iota
	TrafficHistory
	CategoryBrowse
	SitesLinkingIn
)

var (
	AmazonServiceHost = "awis.amazonaws.com"
)

type Request struct {
	// required parameters
	Url           string
	responseGroup string
	action        string

	// internal
	requestType RequestType
	// other parameters
	parameters    map[string]string
	compiledQuery string
}

// Add custom parameter to the request
func (r *Request) SetParam(key, value string) *Request {
	r.parameters[strings.Title(key)] = value
	return r
}

// Signs the query and fills required params.
func (r *Request) compileQuery(AWSKey, secretKey string) {
	preparedParameters := map[string]string{}

	for k, v := range r.parameters {
		preparedParameters[k] = v
	}

	preparedParameters["Url"] = r.Url
	preparedParameters["Action"] = r.action
	preparedParameters["ResponseGroup"] = r.responseGroup
	preparedParameters["AWSAccessKeyId"] = AWSKey
	preparedParameters["SignatureMethod"] = SignatureMethod
	preparedParameters["SignatureVersion"] = SignatureVersion
	preparedParameters["Timestamp"] = time.Now().Format(time.RFC3339)

	query := r.buildQueryString(preparedParameters)

	r.compiledQuery = query + "&Signature=" + sign(AmazonServiceHost, secretKey, query)
}

// Builds query from parameters
func (r *Request) buildQueryString(parameters map[string]string) string {
	var (
		rawQuery string
		keys     []string
	)

	// sort by keys
	for k := range parameters {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		rawQuery += fmt.Sprint(k, "=", strings.Replace(url.QueryEscape(parameters[k]), "+", "%20", -1), "&")
	}

	return strings.TrimRight(rawQuery, "&")
}

func (r *Request) Execute(AWSKey, secretKey string) (res []byte, err error) {
	r.compileQuery(AWSKey, secretKey)

	requestURL := "https://" + AmazonServiceHost + "/?" + r.compiledQuery
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	res, err = ioutil.ReadAll(resp.Body)

	return
}
