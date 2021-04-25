package wxpay

import (
	"encoding/xml"
	"github.com/antchfx/xmlquery"
	"strconv"
	"strings"
)

type RefundQueryRequest struct {
	XMLName       xml.Name `xml:"xml"`
	TransactionId string   `xml:"transaction_id"`
	OutTradeNo    string   `xml:"out_trade_no"`
	OutRefundNo   string   `xml:"out_refund_no"`
	RefundId      string   `xml:"refund_id"`
	Offset        string   `xml:"offset"`
	WxPayRequestImpl
}

func (this *RefundQueryRequest) GetUrl() string {
	return "/pay/refundquery"
}

type RefundQueryResponse struct {
	XMLName            xml.Name `xml:"xml"`
	TotalRefundCount   string   `xml:"total_refund_count"`
	TransactionId      string   `xml:"transaction_id"`
	OutTradeNo         string   `xml:"out_trade_no"`
	TotalFee           string   `xml:"total_fee"`
	SettlementTotalFee string   `xml:"settlement_total_fee"`
	FeeType            string   `xml:"fee_type"`
	CashFee            string   `xml:"cash_fee"`
	RefundCount        int      `xml:"refund_count,omitempty"`
	RefundInfo         []RefundInfo
	WxPayResponseImpl
}
type RefundInfo struct {
	OutRefundNo         string
	RefundId            string
	RefundChannel       string
	RefundFee           string
	SettlementRefundFee string
	CouponRefundFee     string
	CouponRefundCount   int
	RefundStatus        string
	RefundAccount       string
	RefundRecvAccout    string
	RefundSuccessTime   string
	CouponRefund        []CouponRefund
}

type CouponRefund struct {
	CouponRefundId  string
	CouponType      string
	CouponRefundFee string
}

func (this *RefundQueryResponse) AfterPropertiesSet() {
	body := this.ResponseBody

	refundCount := this.RefundCount
	if refundCount == 0 {
		return
	}
	doc, err := xmlquery.Parse(strings.NewReader(body))
	if err != nil {
		return
	}
	root := xmlquery.FindOne(doc, "//xml")
	refundInfos := make([]RefundInfo, refundCount)
	for i := 0; i < refundCount; i++ {
		is := strconv.Itoa(i)
		refundInfos[i] = RefundInfo{}
		if v := root.SelectElement("//coupon_refund_count_" + is); v != nil {
			if couponRefundCount, err := strconv.ParseInt(v.InnerText(), 10, 64); err == nil {
				refundInfos[i].CouponRefundCount = int(couponRefundCount)
			}
		}
		if v := root.SelectElement("//out_refund_no_" + is); v != nil {
			refundInfos[i].OutRefundNo = v.InnerText()
		}
		if v := root.SelectElement("//refund_id_" + is); v != nil {
			refundInfos[i].RefundId = v.InnerText()
		}
		if v := root.SelectElement("//refund_channel_" + is); v != nil {
			refundInfos[i].RefundChannel = v.InnerText()
		}
		if v := root.SelectElement("//refund_fee_" + is); v != nil {
			refundInfos[i].RefundFee = v.InnerText()
		}
		if v := root.SelectElement("//settlement_refund_fee_" + is); v != nil {
			refundInfos[i].SettlementRefundFee = v.InnerText()
		}
		if v := root.SelectElement("//coupon_refund_fee_" + is); v != nil {
			refundInfos[i].CouponRefundFee = v.InnerText()
		}
		if v := root.SelectElement("//refund_status_" + is); v != nil {
			refundInfos[i].RefundStatus = v.InnerText()
		}
		if v := root.SelectElement("//refund_account_" + is); v != nil {
			refundInfos[i].RefundAccount = v.InnerText()
		}
		if v := root.SelectElement("//refund_recv_accout_" + is); v != nil {
			refundInfos[i].RefundRecvAccout = v.InnerText()
		}
		if v := root.SelectElement("//refund_success_time_" + is); v != nil {
			refundInfos[i].RefundSuccessTime = v.InnerText()
		}

		couponRefundCount := refundInfos[i].CouponRefundCount
		couponRefunds := make([]CouponRefund, couponRefundCount)
		if couponRefundCount > 0 {
			for j := 0; j < couponRefundCount; j++ {
				js := strconv.Itoa(j)
				suffix := is + "_" + js
				couponRefunds[j] = CouponRefund{}
				if v := root.SelectElement("//coupon_refund_id_" + suffix); v != nil {
					couponRefunds[i].CouponRefundId = v.InnerText()
				}
				if v := root.SelectElement("//coupon_type_" + suffix); v != nil {
					couponRefunds[i].CouponType = v.InnerText()
				}
				if v := root.SelectElement("//coupon_refund_fee_" + suffix); v != nil {
					couponRefunds[i].CouponRefundFee = v.InnerText()
				}
			}
		}
		refundInfos[i].CouponRefund = couponRefunds
	}
	this.RefundInfo = refundInfos
}
