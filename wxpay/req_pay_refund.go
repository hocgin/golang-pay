package wxpay

import (
	"encoding/xml"
	"github.com/antchfx/xmlquery"
	"strconv"
	"strings"
)

type PayRefundRequest struct {
	XMLName       xml.Name `xml:"xml"`
	TransactionId string   `xml:"transaction_id"`
	OutTradeNo    string   `xml:"out_trade_no"`
	OutRefundNo   string   `xml:"out_refund_no"`
	TotalFee      string   `xml:"total_fee"`
	RefundFee     string   `xml:"refund_fee"`
	RefundFeeType string   `xml:"refund_fee_type"`
	RefundDesc    string   `xml:"refund_desc"`
	RefundAccount string   `xml:"refund_account"`
	NotifyUrl     string   `xml:"notify_url"`
	WxPayRequestImpl
}

func (this *PayRefundRequest) AfterPropertiesSet() {
	this.RequestUrl = "/secapi/pay/refund"
}

type PayRefundResponse struct {
	XMLName             xml.Name `xml:"xml"`
	TransactionId       string   `xml:"transaction_id"`
	OutTradeNo          string   `xml:"out_trade_no"`
	OutRefundNo         string   `xml:"out_refund_no"`
	RefundId            string   `xml:"refund_id"`
	RefundFee           string   `xml:"refund_fee"`
	SettlementRefundFee string   `xml:"settlement_refund_fee"`
	TotalFee            string   `xml:"total_fee"`
	SettlementTotalFee  string   `xml:"settlement_total_fee"`
	FeeType             string   `xml:"fee_type"`
	CashFee             string   `xml:"cash_fee"`
	CashFeeType         string   `xml:"cash_fee_type"`
	CashRefundFee       string   `xml:"cash_refund_fee"`
	CouponRefundFee     string   `xml:"coupon_refund_fee"`
	CouponRefundCount   int      `xml:"coupon_refund_count"`
	CouponRefunds       []PayRefundCouponRefund
	WxPayResponseImpl
}

type PayRefundCouponRefund struct {
	CouponType string `xml:"coupon_type"`
	CouponFee  string `xml:"coupon_fee"`
	CouponId   string `xml:"coupon_id"`
}

func (this *PayRefundResponse) AfterPropertiesSet() {
	body := this.ResponseBody

	couponRefundCount := this.CouponRefundCount
	if couponRefundCount == 0 {
		return
	}
	couponRefunds := make([]PayRefundCouponRefund, couponRefundCount)

	doc, err := xmlquery.Parse(strings.NewReader(body))
	if err != nil {
		return
	}
	root := xmlquery.FindOne(doc, "//xml")
	for i := 0; i < couponRefundCount; i++ {
		is := strconv.Itoa(i)
		couponRefunds[i] = PayRefundCouponRefund{}
		if v := root.SelectElement("//coupon_refund_type_" + is); v != nil {
			couponRefunds[i].CouponType = v.InnerText()
		}
		if v := root.SelectElement("//coupon_refund_id_" + is); v != nil {
			couponRefunds[i].CouponId = v.InnerText()
		}
		if v := root.SelectElement("//coupon_refund_fee_" + is); v != nil {
			couponRefunds[i].CouponFee = v.InnerText()
		}
	}
	this.CouponRefunds = couponRefunds
}
