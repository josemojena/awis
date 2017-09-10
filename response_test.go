package awis

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

const sitesLinkingInRawResponse = `
<?xml version="1.0"?>
<aws:SitesLinkingInResponse xmlns:aws="http://alexa.amazonaws.com/doc/2005-10-05/">
    <aws:Response xmlns:aws="http://awis.amazonaws.com/doc/2005-07-11">
        <aws:OperationRequest>
            <aws:RequestId>ca282ec6-2d08-4341-9f1d-50f8c1e3652b
            </aws:RequestId>
        </aws:OperationRequest>
        <aws:SitesLinkingInResult>
            <aws:Alexa>
                <aws:SitesLinkingIn>
                    <aws:Site>
                        <aws:Title>Google</aws:Title>
                        <aws:Url>http://www.google.com:80/Top/Computers/Internet/On_the_Web/Web_Portals/</aws:Url>
                    </aws:Site>
                    <aws:Site>
                        <aws:Title>www.fotolog.com:80/TsR_BkR_TsR</aws:Title>
                        <aws:Url>http://www.fotolog.com:80/TsR_BkR_TsR</aws:Url>
                    </aws:Site>
                </aws:SitesLinkingIn>
            </aws:Alexa>
        </aws:SitesLinkingInResult>
        <aws:ResponseStatus xmlns:aws="http://alexa.amazonaws.com/doc/2005-10-05/">
            <aws:StatusCode>Success</aws:StatusCode>
        </aws:ResponseStatus>
    </aws:Response>
</aws:SitesLinkingInResponse>
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

func TestUnmarshalSitesLinkingInResponse(t *testing.T) {
	resp := SitesLinkingInResponse{}
	err := xml.Unmarshal([]byte(sitesLinkingInRawResponse), &resp)
	if err != nil {
		t.Error(err)
	}

	if resp.Response.ResponseStatus.StatusCode != "Success" {
		t.Error("response status is not success")
	}

	if resp.Response.SitesLinkingInResult.Alexa.SitesLinkingIn.Site[0].Url != "http://www.google.com:80/Top/Computers/Internet/On_the_Web/Web_Portals/" {
		t.Error("URL wasn't gain")
	}

	if resp.Response.SitesLinkingInResult.Alexa.SitesLinkingIn.Site[1].Title != "www.fotolog.com:80/TsR_BkR_TsR" {
		t.Error("first site title wasn't gain")
	}
}
