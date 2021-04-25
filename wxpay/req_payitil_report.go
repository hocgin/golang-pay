package wxpay

import (
    "encoding/xml"
)

type PayitilReportRequest struct {
    XMLName      xml.Name `xml:"xml"`
    InterfaceUrl string   `xml:"interface_url"`
    ExecuteTime  string   `xml:"execute_time"`
    OutTradeNo   string   `xml:"out_trade_no"`
    UserIp       string   `xml:"user_ip"`
    Time         string   `xml:"time"`
    WxPayRequestImpl
}

func (this *PayitilReportRequest) GetUrl() string {
    return "/payitil/report"
}

type PayitilReportResponse struct {
    WxPayResponseImpl
}
