package alipay

type AliPayMessageImpl struct {
	SignType string `json:"sign_type,omitempty"`
	Sign     string `json:"sign,omitempty"`
}

type AliPayMessage interface {
}
