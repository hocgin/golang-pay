package alipay

type TradeFastpayRefundQueryRequest struct {
	AliPayRequestImpl
	BizContent TradeCreateBizContent `json:"biz_content,omitempty"`
}

type TradeFastpayRefundQueryBizContent struct {
	TradeNo      string `json:"trade_no,omitempty"`
	OutTradeNo   string `json:"out_trade_no,omitempty"`
	OutRequestNo string `json:"out_request_no,omitempty"`
}

func (this *TradeFastpayRefundQueryRequest) AfterPropertiesSet() {
	this.AliPayRequestImpl.AfterPropertiesSet()
	this.Method = "alipay.trade.fastpay.refund.query"
}

type TradeFastpayRefundQueryResponse struct {
	OutTradeNo                   string              `json:"out_trade_no,omitempty"`
	TradeNo                      string              `json:"trade_no,omitempty"`
	OutRequestNo                 string              `json:"out_request_no,omitempty"`
	RefundReason                 string              `json:"refund_reason,omitempty"`
	TotalAmount                  float64             `json:"total_amount,omitempty"`
	RefundAmount                 float64             `json:"refund_amount,omitempty"`
	RefundRoyaltys               RefundRoyaltyResult `json:"refund_royaltys,omitempty"`
	GmtRefundPay                 string              `json:"gmt_refund_pay,omitempty"`
	RefundDetailItemList         TradeFundBill       `json:"refund_detail_item_list,omitempty"`
	SendBackFee                  string              `json:"send_back_fee,omitempty"`
	RefundSettlementId           string              `json:"refund_settlement_id,omitempty"`
	PresentRefundBuyerAmount     float64             `json:"present_refund_buyer_amount,omitempty"`
	PresentRefundDiscountAmount  float64             `json:"present_refund_discount_amount,omitempty"`
	PresentRefundMdiscountAmount float64             `json:"present_refund_mdiscount_amount,omitempty"`
	AliPayResponseImpl
}
type TradeFundBill struct {
	FundChannel string  `json:"fund_channel,omitempty"`
	BankCode    string  `json:"bank_code,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	RealAmount  float64 `json:"real_amount,omitempty"`
}
type RefundRoyaltyResult struct {
	RefundAmount  float64 `json:"refund_amount,omitempty"`
	RoyaltyType   string  `json:"royalty_type,omitempty"`
	ResultCode    string  `json:"result_code,omitempty"`
	TransOut      string  `json:"trans_out,omitempty"`
	TransOutEmail string  `json:"trans_out_email,omitempty"`
	TransIn       string  `json:"trans_in,omitempty"`
	TransInEmail  string  `json:"trans_in_email,omitempty"`
}
