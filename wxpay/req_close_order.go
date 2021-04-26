package wxpay

import (
	"encoding/xml"
)

type CloseOrderRequest struct {
	XMLName    xml.Name `xml:"xml"`
	OutTradeNo string   `xml:"out_trade_no"`
	WxPayRequestImpl
}

func (this *CloseOrderRequest) AfterPropertiesSet() {
	this.RequestUrl = "/pay/closeorder"
}

type CloseOrderResponse struct {
	WxPayResponseImpl
}
