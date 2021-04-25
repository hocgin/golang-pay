package wxpay

import (
    "encoding/xml"
)

type CloseOrderRequest struct {
    XMLName    xml.Name `xml:"xml"`
    OutTradeNo string   `xml:"out_trade_no"`
    WxPayRequestImpl
}

func (this *CloseOrderRequest) GetUrl() string {
    return "/pay/closeorder"
}

type CloseOrderResponse struct {
    WxPayResponseImpl
}
