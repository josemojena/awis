package awis

const AmazonServiceURL = "https://awis.amazonaws.com"

func NewUrlInfoRequest(url string, responseGroup string) *Request {
	return &Request{
		Url:           url,
		responseGroup: responseGroup,
		action:        "UrlInfo",
		requestType:   UrlInfo,

		AmazonServiceURL: AmazonServiceURL,

		parameters: map[string]string{},
	}
}

func NewTrafficHistoryRequest(url string) *Request {
	return &Request{
		Url:           url,
		responseGroup: "History",
		action:        "TrafficHistory",
		requestType:   TrafficHistory,

		AmazonServiceURL: AmazonServiceURL,

		parameters: map[string]string{},
	}
}

func NewCategoryBrowseRequest(url string, responseGroup string) *Request {
	return &Request{
		Url:           url,
		responseGroup: responseGroup,
		action:        "CategoryBrowse",
		requestType:   CategoryBrowse,

		AmazonServiceURL: AmazonServiceURL,

		parameters: map[string]string{},
	}
}

func NewSitesLinkingInRequest(url string) *Request {
	return &Request{
		Url:           url,
		responseGroup: "SitesLinkingIn",
		action:        "SitesLinkingIn",
		requestType:   SitesLinkingIn,

		AmazonServiceURL: AmazonServiceURL,

		parameters: map[string]string{},
	}
}
