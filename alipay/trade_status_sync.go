package alipay

type TradeStatusSyncMessage struct {
	GmtCreate      string         `json:"gmt_create,omitempty"`
	Charset        string         `json:"charset,omitempty"`
	SellerEmail    string         `json:"seller_email,omitempty"`
	Subject        string         `json:"subject,omitempty"`
	BuyerId        string         `json:"buyer_id,omitempty"`
	InvoiceAmount  float64        `json:"invoice_amount,string,omitempty"`
	NotifyId       string         `json:"notify_id,omitempty"`
	NotifyType     string         `json:"notify_type,omitempty"`
	TradeStatus    string         `json:"trade_status,omitempty"`
	ReceiptAmount  float64        `json:"receipt_amount,string,omitempty"`
	AppId          string         `json:"app_id,omitempty"`
	BuyerPayAmount float64        `json:"buyer_pay_amount,string,omitempty"`
	SellerId       string         `json:"seller_id,omitempty"`
	GmtPayment     string         `json:"gmt_payment,omitempty"`
	NotifyTime     string         `json:"notify_time,omitempty"`
	Version        string         `json:"version,omitempty"`
	OutTradeNo     string         `json:"out_trade_no,omitempty"`
	TotalAmount    float64        `json:"total_amount,string,omitempty"`
	TradeNo        string         `json:"trade_no,omitempty"`
	AuthAppId      string         `json:"auth_app_id,omitempty"`
	BuyerLogonId   string         `json:"buyer_logon_id,omitempty"`
	PointAmount    float64        `json:"point_amount,string,omitempty"`
	FundBillList   []FundBillList `json:"fund_bill_list,omitempty"`
	AliPayMessageImpl
}

type FundBillList struct {
	Amount      float64 `json:"amount,string,omitempty"`
	FundChannel string  `json:"fundChannel,omitempty"`
}
