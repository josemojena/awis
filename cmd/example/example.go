package main

import (
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

	fmt.Print(string(res))
}
