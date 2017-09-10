package main

import (
	"encoding/xml"
	"fmt"

	"github.com/un000/awis"
)

func main() {
	r := awis.NewTrafficHistoryRequest("youtube.com")

	res, err := r.SetParam("Start", "20170201").
		Execute("you_aws_key", "your_secret_key")
	if err != nil {
		panic(err)
	}

	resp := awis.TrafficHistoryResponse{}
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
