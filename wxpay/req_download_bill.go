package wxpay

import (
    "encoding/xml"
)

type DownloadBillRequest struct {
    XMLName  xml.Name `xml:"xml"`
    BillDate string   `xml:"bill_date"`
    BillType string   `xml:"bill_type"`
    TarType  string   `xml:"tar_type"`
    WxPayRequestImpl
}

func (this *DownloadBillRequest) GetUrl() string {
    return "/pay/downloadbill"
}

type DownloadBillResponse struct {
    WxPayResponseImpl
}
