package amazon_alexa_web_services_api_client

import (
	"encoding/xml"
	"testing"
)

const trafficHistoryRawResponse = `
<?xml version="1.0"?>
<aws:TrafficHistoryResponse xmlns:aws="http://alexa.amazonaws.com/doc/2005-10-05/"><aws:Response xmlns:aws="http://awis.amazonaws.com/doc/2005-07-11"><aws:OperationRequest><aws:RequestId>1efec9bb-c9b4-80e6-d8ca-bada92605f7c</aws:RequestId></aws:OperationRequest><aws:TrafficHistoryResult><aws:Alexa>

  <aws:TrafficHistory>
    <aws:Range>2</aws:Range>
    <aws:Site>youtube.com</aws:Site>
    <aws:Start>2017-02-01</aws:Start>
    <aws:HistoricalData>
      <aws:Data>
        <aws:Date>2017-02-01</aws:Date>
        <aws:PageViews>
          <aws:PerMillion>47900</aws:PerMillion>
          <aws:PerUser>5.20</aws:PerUser>
        </aws:PageViews>
        <aws:Rank>2</aws:Rank>
        <aws:Reach>
          <aws:PerMillion>390900</aws:PerMillion>
        </aws:Reach>
      </aws:Data>
      <aws:Data>
        <aws:Date>2017-02-02</aws:Date>
        <aws:PageViews>
          <aws:PerMillion>48030</aws:PerMillion>
          <aws:PerUser>5.00</aws:PerUser>
        </aws:PageViews>
        <aws:Rank>2</aws:Rank>
        <aws:Reach>
          <aws:PerMillion>383800</aws:PerMillion>
        </aws:Reach>
      </aws:Data>
    </aws:HistoricalData>
  </aws:TrafficHistory>
</aws:Alexa></aws:TrafficHistoryResult><aws:ResponseStatus xmlns:aws="http://alexa.amazonaws.com/doc/2005-10-05/"><aws:StatusCode>Success</aws:StatusCode></aws:ResponseStatus></aws:Response></aws:TrafficHistoryResponse>
`

func TestUnmarshalTrafficHistoryResponse(t *testing.T) {
	resp := TrafficHistoryResponse{}
	err := xml.Unmarshal([]byte(trafficHistoryRawResponse), &resp)
	if err != nil {
		t.Error(err)
	}

	if resp.Response.ResponseStatus.StatusCode != "Success" {
		t.Error("response status is not success")
	}

	if resp.Response.TrafficHistoryResult.Alexa.TrafficHistory.Site != "youtube.com" {
		t.Error("site wasn't gain")
	}

	if resp.Response.TrafficHistoryResult.Alexa.TrafficHistory.Range != 2 {
		t.Error("range wasn't gain")
	}

	if resp.Response.TrafficHistoryResult.Alexa.TrafficHistory.Start != "2017-02-01" {
		t.Error("start wasn't gain")
	}

	if resp.Response.TrafficHistoryResult.Alexa.TrafficHistory.HistoricalData.Data[0].Date != "2017-02-01" {
		t.Error("HistoricalData.date wasn't gain")
	}

	if resp.Response.TrafficHistoryResult.Alexa.TrafficHistory.HistoricalData.Data[0].PageViews.PerUser != 5.20 {
		t.Error("PageViews.PerUser wasn't gain")
	}

	if resp.Response.TrafficHistoryResult.Alexa.TrafficHistory.HistoricalData.Data[0].Reach.PerMillion != 390900 {
		t.Error("Reach.PerMillion wasn't gain")
	}

	if resp.Response.TrafficHistoryResult.Alexa.TrafficHistory.HistoricalData.Data[0].Rank != 2 {
		t.Error("Rank wasn't gain")
	}
}
