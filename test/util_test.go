package test

import (
	"fmt"
	"rongying-sdk-go/rongying"
	"testing"
)

func TestUtilTimestamp(t *testing.T) {
	fmt.Println(rongying.UtilTimestamp())
}

func TestMdt(t *testing.T) {
	fmt.Println(rongying.Md5("123456"))
}

func TestBase64Encode(t *testing.T) {
	fmt.Println(rongying.Base64Encode("123456"))
}

func TestSendPost(t *testing.T) {
	params := map[string]string{
		"Mobile": "17621166911",
	}
	fmt.Println(rongying.SendPost("https://wdapi.yuntongxin.vip/query/seat/detail/v3?Sig=xxxxxxx", params, "xxxxxxxxx"))
}
