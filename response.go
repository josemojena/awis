package awis

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
		ResponseStatus struct {
			StatusCode string
		}
	}
}
