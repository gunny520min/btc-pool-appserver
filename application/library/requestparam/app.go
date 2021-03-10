package requestparam

type Base struct {
	Comm Common `json:"common"`
}

type Common struct {
	Channel   interface{} `json:"channel"`
	Device    string      `json:"device"`
	DeviceIp  string      `json:"deviceip"`
	Language  string      `json:"language"`
	Nonce     string      `json:"nonce"`
	Platform  string      `json:"platform"`
	Sdk       string      `json:"sdk"`
	Timestamp interface{} `json:"timestamp"`
	Timezone  string      `json:"timezone"`
	Token     string      `json:"token"`
	Version   string      `json:"version"`
}
