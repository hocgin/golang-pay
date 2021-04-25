package alipay

import (
	"github.com/hocgin/golang-pay/core/ops"
)

type TradeCancelRequest struct {
	AliPayRequestImpl
	BizContent TradeCancelBizContent `json:"biz_content,omitempty"`
}
type TradeCancelBizContent struct {
	OutTradeNo string `json:"out_trade_no,omitempty"`
	TradeNo    string `json:"trade_no,omitempty"`
}

func (this *TradeCancelRequest) RequestBefore() {
	this.Method = "alipay.trade.cancel"
}

type TradeCancelResponse struct {
	OutTradeNo         string `json:"out_trade_no,omitempty"`
	TradeNo            string `json:"trade_no,omitempty"`
	RetryFlag          string `json:"retry_flag,omitempty"`
	Action             string `json:"action,omitempty"`
	GmtRefundPay       string `json:"gmt_refund_pay,omitempty"`
	RefundSettlementId string `json:"refund_settlement_id,omitempty"`
	ops.PayResponse
}
