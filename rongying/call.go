package rongying

import (
	"encoding/json"
	"fmt"
)

// ClickCall 点击打电话
func (c *Client) ClickCall(caller string, callee string) (*RyCallResponse, error) {
	ApiUrl := c.conf.buildApiUrl(ClickCall)
	requestParams := &ClickCallRequest{
		Caller: caller,
		Callee: callee,
	}
	params, _ := json.Marshal(requestParams)
	res := SendPost(ApiUrl, string(params), c.Authorization)
	response := &RyCallResponse{}
	err := json.Unmarshal(res.Bytes(), &response)
	return response, err
}

func (c *Client) GetCallRecord(sessionId string) (*RyCallRecordResponse, error) {
	ApiUrl := c.conf.buildApiUrl(CallRecord)
	requestParams := &GetCallRecordRequest{
		CallDetail: CallDetailRequest{
			SessionId: sessionId,
		},
	}
	params, _ := json.Marshal(requestParams)
	fmt.Println(string(params))
	res := SendPost(ApiUrl, string(params), c.Authorization)
	fmt.Println(res)
	response := &RyCallRecordResponse{}
	err := json.Unmarshal(res.Bytes(), &response)
	return response, err
}
