package wxpay

import (
    "encoding/xml"
)

type UnifiedOrderRequest struct {
    XMLName        xml.Name `xml:"xml"`
    DeviceInfo     string   `xml:"device_info"`
    Body           string   `xml:"body"`
    Detail         string   `xml:"detail"`
    Attach         string   `xml:"attach"`
    OutTradeNo     string   `xml:"out_trade_no"`
    FeeType        string   `xml:"fee_type"`
    TotalFee       float64  `xml:"total_fee"`
    SpbillCreateIp string   `xml:"spbill_create_ip"`
    TimeStart      string   `xml:"time_start"`
    TimeExpire     string   `xml:"time_expire"`
    GoodsTag       string   `xml:"goods_tag"`
    NotifyUrl      string   `xml:"notify_url"`
    TradeType      string   `xml:"trade_type"`
    ProductId      string   `xml:"product_id"`
    LimitPay       string   `xml:"limit_pay"`
    Openid         string   `xml:"openid"`
    Receipt        string   `xml:"receipt"`
    SceneInfo      string   `xml:"scene_info"`
    WxPayRequestImpl
}

func (this *UnifiedOrderRequest) GetUrl() string {
    return "/pay/unifiedorder"
}

type UnifiedOrderResponse struct {
    XMLName   xml.Name `xml:"xml"`
    TradeType string   `xml:"trade_type"`
    PrepayId  string   `xml:"prepay_id"`
    CodeUrl   string   `xml:"code_url"`
    WxPayResponseImpl
}
