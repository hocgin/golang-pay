package alipay

import (
	"github.com/hocgin/golang-pay/core/ops"
)

type TradeRefundRequest struct {
	AliPayRequestImpl
	BizContent TradeCreateBizContent `json:"biz_content,omitempty"`
}

type TradeRefundBizContent struct {
	OutTradeNo   string  `json:"out_trade_no,omitempty"`
	TradeNo      string  `json:"trade_no,omitempty"`
	RefundAmount float64 `json:"refund_amount,omitempty"`
}

func (this *TradeRefundRequest) RequestBefore() {
	this.Method = "alipay.trade.refund"
}

type TradeRefundResponse struct {
	OutTradeNo                   string            `json:"out_trade_no,omitempty"`
	TradeNo                      string            `json:"trade_no,omitempty"`
	BuyerLogonId                 string            `json:"buyer_logon_id,omitempty"`
	FundChange                   string            `json:"fund_change,omitempty"`
	RefundFee                    float64           `json:"refund_fee,omitempty"`
	RefundCurrency               string            `json:"refund_currency,omitempty"`
	GmtRefundPay                 string            `json:"gmt_refund_pay,omitempty"`
	StoreName                    string            `json:"store_name,omitempty"`
	BuyerUserId                  string            `json:"buyer_user_id,omitempty"`
	RefundSettlementId           string            `json:"refund_settlement_id,omitempty"`
	PresentRefundBuyerAmount     float64           `json:"present_refund_buyer_amount,omitempty"`
	PresentRefundDiscountAmount  float64           `json:"present_refund_discount_amount,omitempty"`
	PresentRefundMdiscountAmount float64           `json:"present_refund_mdiscount_amount,omitempty"`
	PresetPayToolInfo            PresetPayToolInfo `json:"refund_preset_paytool_list,omitempty"`
	ops.PayResponse
}

type PresetPayToolInfo struct {
	Amount         []string `json:"amount,omitempty"`
	AssertTypeCode string   `json:"assert_type_code,omitempty"`
}
