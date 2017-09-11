package awis

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRequest_Execute(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, trafficHistoryRawResponse)
	}))
	defer ts.Close()

	r := NewTrafficHistoryRequest("youtube.com")
	r.AmazonServiceURL = ts.URL

	res, err := r.SetParam("Start", "20170201").
		Execute("you_aws_key", "your_secret_key")
	if err != nil {
		t.Error(err)
	}

	if strings.TrimSpace(string(res)) != strings.TrimSpace(string(trafficHistoryRawResponse)) {
		t.Error("wrong response")
	}
}

func TestRequest_compileQuery(t *testing.T) {
	requestURL := "youtube.com"
	AWSKey := "aws_key"
	secretKey := "my_secret_key"

	r := NewCategoryBrowseRequest(requestURL, "respGroup")
	r.SetParam("lol", "123")

	if r.compiledQuery != "" {
		t.Errorf("compiled query must be empty at start: %s", r.compiledQuery)
	}

	r.compileQuery(AWSKey, secretKey)

	if !strings.Contains(r.compiledQuery, AWSKey) ||
		!strings.Contains(r.compiledQuery, requestURL) ||
		!strings.Contains(r.compiledQuery, "Lol=123") {
		t.Errorf("compiled query invalid: %s", r.compiledQuery)
	}
}

func TestRequest_buildQueryString(t *testing.T) {
	r := NewTrafficHistoryRequest("google.com")

	r.
		SetParam("a", "random string").
		SetParam("anotherParam", "123@")

	if query := r.buildQueryString(r.parameters); query != "A=random%20string&AnotherParam=123%40" {
		t.Errorf("built query is different: %s", query)
	}
}

func TestRequest_SetParam(t *testing.T) {
	r := NewUrlInfoRequest("google.com", "group")
	r.
		SetParam("a", "b").
		SetParam("c", "d")

	if len(r.parameters) != 2 {
		t.Errorf("wrong paramaeters length %+v", r.parameters)
	}

	if r.parameters["A"] != "b" || r.parameters["C"] != "d" {
		t.Errorf("wrong parameters values %+v", r.parameters)
	}
}

func TestConstructors(t *testing.T) {
	r := NewUrlInfoRequest("google.com", "group")
	if r.requestType != UrlInfo || r.action != "UrlInfo" {
		t.Error("wrong request type")
	}

	r = NewCategoryBrowseRequest("google.com", "group")
	if r.requestType != CategoryBrowse || r.action != "CategoryBrowse" {
		t.Error("wrong request type")
	}

	r = NewTrafficHistoryRequest("google.com")
	if r.requestType != TrafficHistory || r.action != "TrafficHistory" {
		t.Error("wrong request type")
	}

	r = NewSitesLinkingInRequest("google.com")
	if r.requestType != SitesLinkingIn || r.action != "SitesLinkingIn" {
		t.Error("wrong request type")
	}
}

func TestRequest_constructors(t *testing.T) {
	site := "gmail.com"
	respGroup := "cat"

	r := NewTrafficHistoryRequest(site)
	if r.Url != site || r.responseGroup != "History" {
		t.Error("wrong url or respGroup")
	}

	r = NewSitesLinkingInRequest(site)
	if r.Url != site || r.responseGroup != "SitesLinkingIn" {
		t.Error("wrong url or respGroup")
	}

	r = NewUrlInfoRequest(site, respGroup)
	if r.Url != site || r.responseGroup != respGroup {
		t.Error("wrong url or respGroup")
	}

	r = NewCategoryBrowseRequest(site, respGroup)
	if r.Url != site || r.responseGroup != respGroup {
		t.Error("wrong url or respGroup")
	}
}
