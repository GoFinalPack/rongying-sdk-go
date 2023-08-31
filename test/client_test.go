package test

import (
	"fmt"
	"rongying-sdk-go/rongying"
	"testing"
)

var client *rongying.Client

func TestGetSeat(t *testing.T) {
	fmt.Println(client.GetSeat("17854292113")) // 获取指定席位信息
}

func TestGetAllSeat(t *testing.T) {
	fmt.Println(client.GetAllSeat("1")) // 获取所有席位
}

func TestGetCallRecord(t *testing.T) {
	fmt.Println(client.GetCallRecord("xxxxxxxxxxxxxxxxxxx"))
}

func TestMain(m *testing.M) {
	clientConf := &rongying.RyConfig{
		AccountSid:   "xxxxxxxxxxxxxxxxxxxx",
		AccountToken: "xxxxxxxxxxxxxxxxxxxx",
		AppId:        "xxxxxxxxxxxxxxxxxxxx",
		AppToken:     "xxxxxxxxxxxxxxxxxxxx",
	}
	client = rongying.New(clientConf)
	m.Run()
}
