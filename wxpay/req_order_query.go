package wxpay

import (
    "encoding/xml"
    "github.com/antchfx/xmlquery"
    "strconv"
    "strings"
)

type OrderQueryRequest struct {
    XMLName       xml.Name `xml:"xml"`
    TransactionId string   `xml:"transaction_id"`
    OutTradeNo    string   `xml:"out_trade_no"`
    WxPayRequestImpl
}

func (this *OrderQueryRequest) GetUrl() string {
    return "/pay/orderquery"
}

type OrderQueryResponse struct {
    XMLName            xml.Name `xml:"xml"`
    DeviceInfo         string   `xml:"device_info"`
    Openid             string   `xml:"openid"`
    IsSubscribe        string   `xml:"is_subscribe"`
    TradeType          string   `xml:"trade_type"`
    TradeState         string   `xml:"trade_state"`
    BankType           string   `xml:"bank_type"`
    TotalFee           string   `xml:"total_fee"`
    SettlementTotalFee string   `xml:"settlement_total_fee"`
    FeeType            string   `xml:"fee_type"`
    CashFee            string   `xml:"cash_fee"`
    CashFeeType        string   `xml:"cash_fee_type"`
    CouponFee          string   `xml:"coupon_fee"`
    CouponCount        int      `xml:"coupon_count"`
    TransactionId      string   `xml:"transaction_id"`
    OutTradeNo         string   `xml:"out_trade_no"`
    Attach             string   `xml:"attach"`
    TimeEnd            string   `xml:"time_end"`
    TradeStateDesc     string   `xml:"trade_state_desc"`
    Coupons            []Coupon
    WxPayResponseImpl
}

type Coupon struct {
    CouponType string `xml:"coupon_type"`
    CouponFee  string `xml:"coupon_fee"`
    CouponId   string `xml:"coupon_id"`
}

func (this *OrderQueryResponse) AfterPropertiesSet() {
    body := this.ResponseBody

    couponCount := this.CouponCount
    if couponCount == 0 {
        return
    }
    coupons := make([]Coupon, couponCount)

    doc, err := xmlquery.Parse(strings.NewReader(body))
    if err != nil {
        return
    }
    root := xmlquery.FindOne(doc, "//xml")
    for i := 0; i < couponCount; i++ {
        is := strconv.Itoa(i)
        coupons[i] = Coupon{}
        if v := root.SelectElement("//coupon_type_" + is); v != nil {
            coupons[i].CouponType = v.InnerText()
        }
        if v := root.SelectElement("//coupon_id_" + is); v != nil {
            coupons[i].CouponId = v.InnerText()
        }
        if v := root.SelectElement("//coupon_fee_" + is); v != nil {
            coupons[i].CouponFee = v.InnerText()
        }
    }
    this.Coupons = coupons
}
