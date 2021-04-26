package wxpay

import (
	"encoding/xml"
)

type BatchQueryCommentRequest struct {
	XMLName   xml.Name `xml:"xml"`
	BeginTime string   `xml:"begin_time"`
	EndTime   string   `xml:"end_time"`
	Offset    string   `xml:"offset"`
	Limit     string   `xml:"limit"`
	WxPayRequestImpl
}

func (this *BatchQueryCommentRequest) AfterPropertiesSet() {
	this.RequestUrl = "/billcommentsp/batchquerycomment"
}

type BatchQueryCommentResponse struct {
	WxPayResponseImpl
}
