package alipay

import (
	"github.com/hocgin/golang-pay/core/ops"
)

type TradeAdvanceConsultRequest struct {
	AliPayRequestImpl
	BizContent TradeAdvanceConsultBizContent `json:"biz_content,omitempty"`
}
type TradeAdvanceConsultBizContent struct {
	OutTradeNo   string `json:"out_trade_no,omitempty"`
	AlipayUserId string `json:"alipay_user_id,omitempty"`
	ConsultScene string `json:"consult_scene,omitempty"`
	AgreementNo  string `json:"agreement_no,omitempty"`
}

func (this *TradeAdvanceConsultRequest) RequestBefore() {
	this.Method = "alipay.trade.advance.consult"
}

type TradeAdvanceConsultResponse struct {
	ReferResult             string                 `json:"refer_result,omitempty"`
	WaitRepaymentOrderInfos WaitRepaymentOrderInfo `json:"wait_repayment_order_infos,omitempty"`
	WaitRepaymentAmount     float64                `json:"wait_repayment_amount,omitempty"`
	WaitRepaymentOrderCount string                 `json:"wait_repayment_order_count,omitempty"`
	RiskLevel               string                 `json:"risk_level,omitempty"`
	ResultMessage           string                 `json:"result_message,omitempty"`
	ResultCode              string                 `json:"result_code,omitempty"`
	ops.PayResponse
}

type WaitRepaymentOrderInfo struct {
	AdvanceOrderId      string `json:"advance_order_id,omitempty"`
	AlipayUserId        string `json:"alipay_user_id,omitempty"`
	OrigBizOrderId      string `json:"orig_biz_order_id,omitempty"`
	BizProduct          string `json:"biz_product,omitempty"`
	WaitRepaymentAmount string `json:"wait_repayment_amount,omitempty"`
}
