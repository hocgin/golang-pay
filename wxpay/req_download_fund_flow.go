package wxpay

import (
	"encoding/xml"
)

type DownloadFundFlowRequest struct {
	XMLName     xml.Name `xml:"xml"`
	BillDate    string   `xml:"bill_date"`
	AccountType string   `xml:"account_type"`
	TarType     string   `xml:"tar_type"`
	WxPayRequestImpl
}

func (this *DownloadFundFlowRequest) AfterPropertiesSet() {
	this.RequestUrl = "/pay/downloadfundflow"
}

type DownloadFundFlowResponse struct {
	WxPayResponseImpl
}
