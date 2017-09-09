package main

import (
	"encoding/xml"
	"fmt"
	c "github.com/un000/amazon-alexa-web-services-api-client"
)

func main() {
	r := c.NewTrafficHistoryRequest("youtube.com")

	res, err := r.SetParam("Start", "20170201").
		Execute("you_aws_key", "your_secret_key")
	if err != nil {
		panic(err)
	}

	resp := c.TrafficHistoryResponse{}
	err = xml.Unmarshal(res, &resp)
	if err != nil {
		panic(err)
	}

	if resp.Response.ResponseStatus.StatusCode == "Success" {
		fmt.Print("success!")
	} else {
		fmt.Print("bad status code: ", resp.Response.ResponseStatus.StatusCode)
	}
}
