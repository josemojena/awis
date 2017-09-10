package awis

import (
	"testing"
)

func TestSign(t *testing.T) {
	query := "AWSAccessKeyId=rrr&Action=TrafficHistory&SignatureMethod=HmacSHA256&SignatureVersion=2&Start=20170201&Timestamp=2017-09-09T23%3A43%3A10%2B03%3A00&Url=youtube.com"
	hash := sign("awis.amazonaws.com", "secret", query)

	if hash != "YIMcjBPKZR92jibhFShiEwPrryn0QTWl2R4Yve%2BrMa4%3D" {
		t.Errorf("hash not equals %s", hash)
	}
}
