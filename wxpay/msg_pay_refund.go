package wxpay

import (
	"encoding/xml"
	"github.com/hocgin/golang-pay/core"
)

type PayRefundMessage struct {
	Appid    string `xml:"appid,omitempty"`
	MchId    string `xml:"mch_id,omitempty"`
	NonceStr string `xml:"nonce_str,omitempty"`
	ReqInfo  ReqInfo
	WxPayMessageImpl
}

type ReqInfo struct {
	TransactionId       string `xml:"transaction_id,string,omitempty"`
	OutTradeNo          string `xml:"out_trade_no,omitempty"`
	OutRefundNo         string `xml:"out_refund_no,omitempty"`
	TotalFee            string `xml:"total_fee,omitempty"`
	SettlementTotalFee  string `xml:"settlement_total_fee,omitempty"`
	RefundFee           string `xml:"refund_fee,omitempty"`
	SettlementRefundFee string `xml:"settlement_refund_fee,omitempty"`
	RefundStatus        string `xml:"refund_status,omitempty"`
	SuccessTime         string `xml:"success_time,omitempty"`
	RefundRecvAccout    string `xml:"refund_recv_accout,omitempty"`
	RefundAccount       string `xml:"refund_account,omitempty"`
	RefundRequestSource string `xml:"refund_request_source,omitempty"`
}

func (this *PayRefundMessage) AfterPropertiesSet() {
	service := this.PayService.(WxPayPaymentService)
	key := core.MD5(service.ConfigStorage.Key)
	value := this.getXmlValue("xml/req_info")
	reqInfoString := core.AesECBDecrypt(value, key)
	reqInfo := ReqInfo{}
	_ = xml.Unmarshal([]byte(reqInfoString), &reqInfo)
	this.ReqInfo = reqInfo
}
