package alipay

type TradeAppPayRequest struct {
	AliPayRequestImpl
	BizContent TradeAppPayBizContent `json:"biz_content,omitempty"`
}
type TradeAppPayBizContent struct {
	OutTradeNo  string  `json:"out_trade_no,omitempty"`
	Subject     string  `json:"subject,omitempty"`
	TotalAmount float64 `json:"total_amount,omitempty"`
}

func (this *TradeAppPayRequest) AfterPropertiesSet() {
	this.AliPayRequestImpl.AfterPropertiesSet()
	this.Method = "alipay.trade.app.pay"
}

type TradeAppPayResponse struct {
	AliPayResponseImpl
}
