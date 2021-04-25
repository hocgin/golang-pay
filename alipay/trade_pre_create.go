package alipay

type TradePreCreateRequest struct {
	AliPayRequestImpl
	BizContent TradePreCreateBizContent `json:"biz_content,omitempty"`
}

type TradePreCreateBizContent struct {
	OutTradeNo  string  `json:"out_trade_no,omitempty"`
	TotalAmount float64 `json:"total_amount,omitempty"`
	Subject     string  `json:"subject,omitempty"`
}

func (this *TradePreCreateRequest) RequestBefore() {
	this.Method = "alipay.trade.precreate"
}

type TradePreCreateResponse struct {
	OutTradeNo string `json:"out_trade_no,omitempty"`
	QrCode     string `json:"qr_code,omitempty"`
	AliPayResponseImpl
}
