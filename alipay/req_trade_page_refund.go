package alipay

type TradePageRefundRequest struct {
	AliPayRequestImpl
	BizContent TradePageRefundBizContent `json:"biz_content,omitempty"`
}

type TradePageRefundBizContent struct {
	OutTradeNo   string  `json:"out_trade_no,omitempty"`
	TradeNo      string  `json:"trade_no,omitempty"`
	OutRequestNo string  `json:"out_request_no,omitempty"`
	RefundAmount float64 `json:"refund_amount,omitempty"`
}

func (this *TradePageRefundRequest) AfterPropertiesSet() {
	this.AliPayRequestImpl.AfterPropertiesSet()
	this.Method = "alipay.trade.page.refund"
}

type TradePageRefundResponse struct {
	OutTradeNo   string `json:"out_trade_no,omitempty"`
	TradeNo      string `json:"trade_no,omitempty"`
	OutRequestNo string `json:"out_request_no,omitempty"`
	RefundAmount string `json:"refund_amount,omitempty"`
	AliPayResponseImpl
}
