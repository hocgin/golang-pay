package alipay

import (
	"github.com/hocgin/golang-pay/core/ops"
)

type TradePagePayRequest struct {
	AliPayRequestImpl
	BizContent TradeCreateBizContent `json:"biz_content,omitempty"`
}

type TradePagePayBizContent struct {
	OutTradeNo  string  `json:"out_trade_no,omitempty"`
	ProductCode string  `json:"product_code,omitempty"`
	TotalAmount float64 `json:"total_amount,omitempty"`
	Subject     string  `json:"subject,omitempty"`
}

func (this *TradePagePayRequest) RequestBefore() {
	this.Method = "alipay.trade.page.pay"
}

type TradePagePayResponse struct {
	ops.PayResponse
}
