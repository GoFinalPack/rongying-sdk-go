package rongying

// GetSeatRequest 获取指定坐席的 请求结构体
type GetSeatRequest struct {
	Mobile string
}

// ClickCallRequest 点击打电话
type ClickCallRequest struct {
	Caller string
	Callee string
}
type CallDetailRequest struct {
	SessionId string
}

// GetCallRecordRequest 获取通话记录
type GetCallRecordRequest struct {
	CallDetail CallDetailRequest
}
