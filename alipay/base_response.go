package alipay

import "github.com/hocgin/golang-pay/core/ops"

type AliPayResponseImpl struct {
	Code    string `json:"code,omitempty"`
	Msg     string `json:"msg,omitempty"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
	ops.PayResponse
}

type AliPayResponse interface {
}
