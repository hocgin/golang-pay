package alipay

import (
	"github.com/hocgin/golang-pay/core/ops"
)

type TradeWapPayRequest struct {
	AliPayRequestImpl
	BizContent TradeCreateBizContent `json:"biz_content,omitempty"`
}

type TradeWapPayBizContent struct {
	OutTradeNo  string  `json:"out_trade_no,omitempty"`
	ProductCode string  `json:"product_code,omitempty"`
	QuitUrl     string  `json:"quit_url,omitempty"`
	TotalAmount float64 `json:"total_amount,omitempty"`
	Subject     string  `json:"subject,omitempty"`
}

func (this *TradeWapPayRequest) RequestBefore() {
	this.Method = "alipay.trade.wap.pay"
}

type TradeWapPayResponse struct {
	ops.PayResponse
}