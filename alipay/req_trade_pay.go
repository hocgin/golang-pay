package alipay

type TradePayRequest struct {
	AliPayRequestImpl
	BizContent TradePayBizContent `json:"biz_content,omitempty"`
}
type TradePayBizContent struct {
	OutTradeNo string `json:"out_trade_no,omitempty"`
	Scene      string `json:"scene,omitempty"`
	AuthCode   string `json:"auth_code,omitempty"`
	Subject    string `json:"subject,omitempty"`
}

func (this *TradePayRequest) AfterPropertiesSet() {
	this.AliPayRequestImpl.AfterPropertiesSet()
	this.Method = "alipay.trade.pay"
}

type TradePayResponse struct {
	OutTradeNo          string        `json:"out_trade_no,omitempty"`
	TradeNo             string        `json:"trade_no,omitempty"`
	BuyerLogonId        string        `json:"buyer_logon_id,omitempty"`
	SettleAmount        float64       `json:"settle_amount,omitempty"`
	PayCurrency         string        `json:"pay_currency,omitempty"`
	PayAmount           float64       `json:"pay_amount,omitempty"`
	SettleTransRate     string        `json:"settle_trans_rate,omitempty"`
	TransPayRate        string        `json:"trans_pay_rate,omitempty"`
	TotalAmount         float64       `json:"total_amount,omitempty"`
	TransCurrency       string        `json:"trans_currency,omitempty"`
	SettleCurrency      string        `json:"settle_currency,omitempty"`
	ReceiptAmount       float64       `json:"receipt_amount,omitempty"`
	BuyerPayAmount      float64       `json:"buyer_pay_amount,omitempty"`
	PointAmount         float64       `json:"point_amount,omitempty"`
	InvoiceAmount       float64       `json:"invoice_amount,omitempty"`
	GmtPayment          string        `json:"gmt_payment,omitempty"`
	FundBillList        TradeFundBill `json:"fund_bill_list,omitempty"`
	CardBalance         string        `json:"card_balance,omitempty"`
	StoreName           string        `json:"store_name,omitempty"`
	BuyerUserId         string        `json:"buyer_user_id,omitempty"`
	DiscountGoodsDetail string        `json:"discount_goods_detail,omitempty"`
	VoucherDetailList   VoucherDetail `json:"voucher_detail_list,omitempty"`
	AdvanceAmount       string        `json:"advance_amount,omitempty"`
	AuthTradePayMode    string        `json:"auth_trade_pay_mode,omitempty"`
	ChargeAmount        string        `json:"charge_amount,omitempty"`
	ChargeFlags         string        `json:"charge_flags,omitempty"`
	SettlementId        string        `json:"settlement_id,omitempty"`
	BusinessParams      string        `json:"business_params,omitempty"`
	BuyerUserType       string        `json:"buyer_user_type,omitempty"`
	MdiscountAmount     string        `json:"mdiscount_amount,omitempty"`
	DiscountAmount      string        `json:"discount_amount,omitempty"`
	BuyerUserName       string        `json:"buyer_user_name,omitempty"`
	AliPayResponseImpl
}
type VoucherDetail struct {
	Id                         string `json:"id,omitempty"`
	Name                       string `json:"name,omitempty"`
	Type                       string `json:"type,omitempty"`
	Amount                     string `json:"amount,omitempty"`
	MerchantContribute         string `json:"merchant_contribute,omitempty"`
	OtherContribute            string `json:"other_contribute,omitempty"`
	Memo                       string `json:"memo,omitempty"`
	TemplateId                 string `json:"template_id,omitempty"`
	PurchaseBuyerContribute    string `json:"purchase_buyer_contribute,omitempty"`
	PurchaseMerchantContribute string `json:"purchase_merchant_contribute,omitempty"`
	PurchaseAntContribute      string `json:"purchase_ant_contribute,omitempty"`
}
