package amazon_alexa_web_services_api_client

type TrafficHistoryResponse struct {
	Response struct {
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
								PerMillion int
							}
							PageViews struct {
								PerMillion int
								PerUser    string
							}
						}
					}
				}
			}
		}
		ResponseStatus struct {
			StatusCode string
		}
	}
}
