package awis

func NewUrlInfoRequest(url string, responseGroup string) *Request {
	return &Request{
		Url:           url,
		responseGroup: responseGroup,
		action:        "UrlInfo",
		requestType:   UrlInfo,

		parameters: map[string]string{},
	}
}

func NewTrafficHistoryRequest(url string) *Request {
	return &Request{
		Url:           url,
		responseGroup: "History",
		action:        "TrafficHistory",
		requestType:   TrafficHistory,

		parameters: map[string]string{},
	}
}

func NewCategoryBrowseRequest(url string, responseGroup string) *Request {
	return &Request{
		Url:           url,
		responseGroup: responseGroup,
		action:        "CategoryBrowse",
		requestType:   CategoryBrowse,

		parameters: map[string]string{},
	}
}

func NewSitesLinkingInRequest(url string) *Request {
	return &Request{
		Url:           url,
		responseGroup: "SitesLinkingIn",
		action:        "SitesLinkingIn",
		requestType:   SitesLinkingIn,

		parameters: map[string]string{},
	}
}
