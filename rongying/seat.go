package rongying

import (
	"encoding/json"
	"fmt"
)

// GetSeat 获取指定坐席
func (c *Client) GetSeat(mobile string) (*RySeatListResponse, error) {
	ApiUrl := c.conf.buildApiUrl(Seat)
	requestParam := &GetSeatRequest{
		Mobile: mobile,
	}
	params, _ := json.Marshal(requestParam)
	fmt.Println(string(params))
	res := SendPost(ApiUrl, string(params), c.Authorization)
	response := &RySeatListResponse{}
	err := json.Unmarshal(res.Bytes(), &response)
	return response, err
}

func (c *Client) GetAllSeat(page string) (*RySeatListResponse, error) {
	ApiUrl := c.conf.buildApiUrl(SeatList)
	requestParam := map[string]string{
		"Page": page,
	}
	params, _ := json.Marshal(requestParam)
	fmt.Println(string(params))
	res := SendPost(ApiUrl, string(params), c.Authorization)
	response := &RySeatListResponse{}
	err := json.Unmarshal(res.Bytes(), &response)
	return response, err
}
