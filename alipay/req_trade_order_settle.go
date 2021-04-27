package alipay

type TradeOrderSettleRequest struct {
	AliPayRequestImpl
	BizContent TradeOrderSettleBizContent `json:"biz_content,omitempty"`
}

type TradeOrderSettleBizContent struct {
	OutRequestNo      string                         `json:"out_request_no,omitempty"`
	TradeNo           string                         `json:"trade_no,omitempty"`
	RoyaltyParameters []OpenApiRoyaltyDetailInfoPojo `json:"royalty_parameters,omitempty"`
}
type OpenApiRoyaltyDetailInfoPojo struct {
	TransIn string `json:"trans_in,omitempty"`
}

func (this *TradeOrderSettleRequest) AfterPropertiesSet() {
	this.AliPayRequestImpl.AfterPropertiesSet()
	this.Method = "alipay.trade.order.settle"
}

type TradeOrderSettleResponse struct {
	TradeNo string `json:"trade_no,omitempty"`
	AliPayResponseImpl
}
