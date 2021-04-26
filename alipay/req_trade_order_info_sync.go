package alipay

type TradeOrderInfoSyncRequest struct {
	AliPayRequestImpl
	BizContent TradeCreateBizContent `json:"biz_content,omitempty"`
}

type TradeOrderInfoSyncBizContent struct {
	TradeNo      string `json:"trade_no,omitempty"`
	OutRequestNo string `json:"out_request_no,omitempty"`
	BizType      string `json:"biz_type,omitempty"`
}

func (this *TradeOrderInfoSyncRequest) AfterPropertiesSet() {
	this.AliPayRequestImpl.AfterPropertiesSet()
	this.Method = "alipay.trade.orderinfo.sync"
}

type TradeOrderInfoSyncResponse struct {
	OutTradeNo  string `json:"out_trade_no,omitempty"`
	TradeNo     string `json:"trade_no,omitempty"`
	BuyerUserId string `json:"buyer_user_id,omitempty"`
	AliPayResponseImpl
}
