package wxpay

import (
	"strconv"
)

type UnifiedOrderMessage struct {
	Appid              string  `xml:"appid,omitempty"`
	MchId              string  `xml:"mch_id,omitempty"`
	NonceStr           string  `xml:"nonce_str,omitempty"`
	Sign               string  `xml:"sign,omitempty"`
	SignType           string  `xml:"sign_type,omitempty"`
	ResultCode         string  `xml:"result_code,omitempty"`
	ErrCode            string  `xml:"err_code,omitempty"`
	ErrCodeDes         string  `xml:"err_code_des,omitempty"`
	DeviceInfo         string  `xml:"device_info,omitempty"`
	Openid             string  `xml:"openid,omitempty"`
	IsSubscribe        string  `xml:"is_subscribe,omitempty"`
	TradeType          string  `xml:"trade_type,omitempty"`
	BankType           string  `xml:"bank_type,omitempty"`
	TotalFee           float64 `xml:"total_fee,omitempty"`
	SettlementTotalFee string  `xml:"settlement_total_fee,omitempty"`
	FeeType            string  `xml:"fee_type,omitempty"`
	CashFee            float64 `xml:"cash_fee,omitempty"`
	CashFeeType        string  `xml:"cash_fee_type,omitempty"`
	CouponFee          string  `xml:"coupon_fee,omitempty"`
	CouponCount        int     `xml:"coupon_count,omitempty"`
	TransactionId      string  `xml:"transaction_id,omitempty"`
	OutTradeNo         string  `xml:"out_trade_no,omitempty"`
	Attach             string  `xml:"attach,omitempty"`
	TimeEnd            string  `xml:"time_end,omitempty"`

	Coupons []Coupon
	WxPayMessageImpl
}

func (this *UnifiedOrderMessage) AfterPropertiesSet() {
	coupons := make([]Coupon, this.CouponCount)
	if this.CouponCount > 0 {
		root, err := this.getXmlDoc("xml")
		if err != nil {
			panic(err)
		}
		for i := 0; i < this.CouponCount; i++ {
			suffix := strconv.Itoa(i)
			coupons[i] = Coupon{}
			if v := root.SelectElement("//coupon_id_" + suffix); v != nil {
				coupons[i].CouponId = v.InnerText()
			}
			if v := root.SelectElement("//coupon_type_" + suffix); v != nil {
				coupons[i].CouponType = v.InnerText()
			}
			if v := root.SelectElement("//coupon_fee_" + suffix); v != nil {
				coupons[i].CouponFee = v.InnerText()
			}
		}
	}
	this.Coupons = coupons
}
