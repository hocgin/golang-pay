package alipay

type TradeCloseRequest struct {
	AliPayRequestImpl
	BizContent TradeCloseBizContent `json:"biz_content,omitempty"`
}
type TradeCloseBizContent struct {
	OutTradeNo string `json:"out_trade_no,omitempty"`
	TradeNo    string `json:"trade_no,omitempty"`
}

func (this *TradeCloseRequest) AfterPropertiesSet() {
	this.AliPayRequestImpl.AfterPropertiesSet()
	this.Method = "alipay.trade.close"
}

type TradeCloseResponse struct {
	OutTradeNo string `json:"out_trade_no,omitempty"`
	TradeNo    string `json:"trade_no,omitempty"`
	AliPayResponseImpl
}
