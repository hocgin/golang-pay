package alipay

type TradeQueryRequest struct {
	AliPayRequestImpl
	BizContent TradeQueryBizContent `json:"biz_content,omitempty"`
}

type TradeQueryBizContent struct {
	OutTradeNo string `json:"out_trade_no,omitempty"`
	TradeNo    string `json:"trade_no,omitempty"`
}

func (this *TradeQueryRequest) AfterPropertiesSet() {
	this.AliPayRequestImpl.AfterPropertiesSet()
	this.Method = "alipay.trade.query"
}

type TradeQueryResponse struct {
	OutTradeNo          string        `json:"out_trade_no,omitempty"`
	TradeNo             string        `json:"trade_no,omitempty"`
	BuyerLogonId        string        `json:"buyer_logon_id,omitempty"`
	TradeStatus         string        `json:"trade_status,omitempty"`
	TotalAmount         float64       `json:"total_amount,omitempty"`
	TransCurrency       string        `json:"trans_currency,omitempty"`
	SettleCurrency      string        `json:"settle_currency,omitempty"`
	SettleAmount        float64       `json:"settle_amount,omitempty"`
	PayCurrency         string        `json:"pay_currency,omitempty"`
	PayAmount           float64       `json:"pay_amount,omitempty"`
	SettleTransRate     string        `json:"settle_trans_rate,omitempty"`
	TransPayRate        string        `json:"trans_pay_rate,omitempty"`
	BuyerPayAmount      float64       `json:"buyer_pay_amount,omitempty"`
	PointAmount         float64       `json:"point_amount,omitempty"`
	InvoiceAmount       float64       `json:"invoice_amount,omitempty"`
	SendPayDate         string        `json:"send_pay_date,omitempty"`
	StoreId             string        `json:"store_id,omitempty"`
	TerminalId          string        `json:"terminal_id,omitempty"`
	TradeFundBill       TradeFundBill `json:"fund_bill_list,omitempty"`
	ReceiptAmount       float64       `json:"receipt_amount,omitempty"`
	StoreName           string        `json:"store_name,omitempty"`
	BuyerUserId         string        `json:"buyer_user_id,omitempty"`
	ChargeAmount        string        `json:"charge_amount,omitempty"`
	ChargeFlags         string        `json:"charge_flags,omitempty"`
	SettlementId        string        `json:"settlement_id,omitempty"`
	TradeSettleInfo     string        `json:"trade_settle_info,omitempty"`
	AuthTradePayMode    string        `json:"auth_trade_pay_mode,omitempty"`
	BuyerUserType       string        `json:"buyer_user_type,omitempty"`
	MdiscountAmount     string        `json:"mdiscount_amount,omitempty"`
	DiscountAmount      string        `json:"discount_amount,omitempty"`
	BuyerUserName       string        `json:"buyer_user_name,omitempty"`
	Subject             string        `json:"subject,omitempty"`
	Body                string        `json:"body,omitempty"`
	AlipaySubMerchantId string        `json:"alipay_sub_merchant_id,omitempty"`
	ExtInfos            string        `json:"ext_infos,omitempty"`
	AliPayResponseImpl
}

type TradeSettleInfo struct {
	OperationType     string  `json:"operation_type,omitempty"`
	OperationSerialNo string  `json:"operation_serial_no,omitempty"`
	OperationDt       string  `json:"operation_dt,omitempty"`
	TransOut          string  `json:"trans_out,omitempty"`
	TransIn           string  `json:"trans_in,omitempty"`
	Amount            float64 `json:"amount,omitempty"`
}
