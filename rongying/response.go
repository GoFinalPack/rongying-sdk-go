package rongying

type RyResponse struct {
	Flag int
	Msg  string
}
type RySeat struct {
	SeatAccount string
	Passwd      string `json:"Passwd"`
	State       int    `json:"State"`
	Mobile      string `json:"Mobile"`
	IsVail      int    `json:"IsVail"`
	SeatName    string `json:"SeatName"`
	IsExamine   int    `json:"IsExamine"`
}
type RySeatListResponse struct {
	Flag int
	Msg  string
	Data []RySeat
}

type RyCallResponse struct {
	Flag          int
	Msg           string
	VirtualNumber string
	SessionId     string
}

type RyCallRecordResponseData struct {
	maxid                 int
	sessionId             string
	bindNum               string
	calleeNum             string
	fwdDstNum             string
	fwdDisplayNum         string
	fwdStartTime          string
	fwdAlertingTime       string
	fwdAnswerTime         string
	callEndTime           string
	failTime              string
	callOutStartTime      string
	callOutAlertingTime   string
	callOutAnswerTime     string
	billsec               int
	recordFlag            int
	recordStartTime       string
	recordFileDownloadUrl string
	fwdUnaswRsn           string
	ulFailReason          string
	direction             string
}

type RyCallRecordResponse struct {
	Flag int
	Msg  string
	Data map[string]RyCallRecordResponseData
}
