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
	fmt.Println(rongying.SendPost("https://wdapi.yuntongxin.vip/query/seat/detail/v3?Sig=5AE42666B0BB03A982DBC06D41175135", params, "Njc2NzYxOGQzYzc1NGI1M2E3MTk1NjQ5YTdiYTAzM2Q4MDRjYjYzNWMwMjE0ODc2YWIxODJlOWJhM2FhOTgxYjIwMjMwODMwMTYwNDM4"))
}
