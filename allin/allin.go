package allin

type AllInPayService interface {
	DoPay(trade TradeInfo) (*PayResult, error)
	DoRefund(RefundInfo) error
	DoQueryPay(outTradeNo string) (*QueryPayResult, error)
	DoQueryRefund(outTradeNo string) (*QueryRefundResult, error)
	DoPayMessage(queryStrings string) (*PayMessageResult, error)
	DoRefundMessage(queryStrings string) (*RefundMessageResult, error)
}
