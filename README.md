rongying-sdk-go 

---

####  融营 电话呼叫 SDK golang 版本

####  结构

```azure
.
├── README.md
├── go.mod
├── go.sum
├── rongying
│   ├── call.go      -- 呼叫相关
│   ├── client.go  
│   ├── config.go    -- 配置签名相关
│   ├── request.go   -- 所有的请求参数相关
│   ├── response.go  -- 所有的相应相关
│   ├── seat.go      -- 客服相关
│   └── util.go      -- 常用函数相关
├── rongying.go
└── test                   -- 测试相关
    ├── client_test.go
    └── util_test.go

```

####  使用方式

```golang
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
	fmt.Println(client.GetAllSeat("1"))   // 获取所有席位
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


```

