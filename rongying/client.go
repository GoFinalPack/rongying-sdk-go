package rongying

type Client struct {
	conf          *RyConfig
	Sig           string
	Authorization string
}

func New(clientConf *RyConfig) *Client {
	Sig, Authorization := clientConf.initAuth()
	return &Client{
		conf:          clientConf,
		Sig:           Sig,
		Authorization: Authorization,
	}
}
