package awis

type OperationRequest struct {
	RequestId string
}

type ResponseStatus struct {
	StatusCode string
}

type SitesLinkingInResponse struct {
	Response struct {
		OperationRequest     OperationRequest
		SitesLinkingInResult struct {
			Alexa struct {
				SitesLinkingIn struct {
					Site []struct {
						Title string
						Url   string
					}
				}
			}
		}
		ResponseStatus ResponseStatus
	}
}

type TrafficHistoryResponse struct {
	Response struct {
		OperationRequest     OperationRequest
		TrafficHistoryResult struct {
			Alexa struct {
				TrafficHistory struct {
					Range          int
					Site           string
					Start          string
					HistoricalData struct {
						Data []struct {
							Date  string
							Rank  int
							Reach struct {
								PerMillion float64
							}
							PageViews struct {
								PerMillion float64
								PerUser    float64
							}
						}
					}
				}
			}
		}
		ResponseStatus ResponseStatus
	}
}
