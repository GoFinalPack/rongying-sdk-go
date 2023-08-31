package rongying

const (
	ResetUrl   = "https://wdapi.yuntongxin.vip"
	Seat       = "/query/seat/detail/v3"
	SeatList   = "/query/account/seatlist/v3"
	AppInfo    = "/rest/query/account/v1"
	ClickCall  = "/rest/click/call/event/v2"
	CallRecord = "/20200306/rest/click/call/record/v1"
)

type RyConfig struct {
	AccountSid   string
	AccountToken string
	AppId        string
	AppToken     string
}

func (rc *RyConfig) initAuth() (string, string) {
	timestamp := UtilTimestamp()
	sign := rc.AccountSid + ":" + rc.AppId + ":" + timestamp
	auth := rc.AppId + ":" + rc.AppToken + ":" + timestamp
	return Md5(sign), Base64Encode(auth)
}

func (rc *RyConfig) initAppAuth() (string, string) {
	timestamp := UtilTimestamp()
	sign := rc.AccountSid + ":" + rc.AppToken + ":" + timestamp
	auth := rc.AccountSid + ":" + rc.AccountToken + ":" + timestamp
	return Md5(sign), Base64Encode(auth)
}

func (rc *RyConfig) buildApiUrl(url string) string {
	Sig, _ := rc.initAuth()
	return ResetUrl + url + "?Sig=" + Sig
}
