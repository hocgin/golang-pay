package alipay

import (
	"github.com/hocgin/golang-pay/core/ops"
)

type TradeCreateRequest struct {
	AliPayRequestImpl
	BizContent TradeCreateBizContent `json:"biz_content,omitempty"`
}

type TradeCreateBizContent struct {
	OutTradeNo  string  `json:"out_trade_no,omitempty"`
	TotalAmount float64 `json:"total_amount,omitempty"`
	BuyerId     string  `json:"buyer_id,omitempty"`
	Subject     string  `json:"subject,omitempty"`
}

func (this *TradeCreateRequest) RequestBefore() {
	this.Method = "alipay.trade.create"
}

type TradeCreateResponse struct {
	OutTradeNo string `json:"out_trade_no,omitempty"`
	TradeNo    string `json:"trade_no,omitempty"`
	ops.PayResponse
}
