package wxpay

import (
	"errors"
	"github.com/hocgin/golang-pay/allin"
	"github.com/hocgin/golang-pay/core"
	"time"
)

// DoPay 去支付
func (this WxPayPaymentService) DoPay(trade allin.TradeInfo) (*allin.PayResult, error) {
	result := new(allin.PayResult)
	result.PayWay = trade.PayWay
	var tradeType string
	switch trade.PayWay {
	case allin.WxPayApp:
		tradeType = "APP"
	case allin.WxPayJsApi:
		tradeType = "JSAPI"
	case allin.WxPayNative:
		tradeType = "NATIVE"
	default:
		return nil, errors.New("wxpay not support the pay way")
	}
	rep, err := this.UnifiedOrder(UnifiedOrderRequest{OutTradeNo: trade.OutTradeNo,
		TotalFee:       trade.TotalAmount,
		TradeType:      tradeType,
		Body:           trade.Subject,
		Openid:         trade.Openid,
		NotifyUrl:      trade.NotifyUrl,
		SpbillCreateIp: "192.168.1.1"})
	result.Response = rep
	return result, err
}

// DoRefund 去退款
func (this WxPayPaymentService) DoRefund(refund allin.RefundInfo) error {
	_, err := this.PayRefund(PayRefundRequest{OutTradeNo: refund.OutTradeNo, RefundFee: refund.RefundAmount})
	return err
}

// DoQueryPay 查询支付
func (this WxPayPaymentService) DoQueryPay(outTradeNo string) (*allin.QueryPayResult, error) {
	_, err := this.OrderQuery(OrderQueryRequest{OutTradeNo: outTradeNo})
	return nil, err
}

// DoQueryRefund 查询退款
func (this WxPayPaymentService) DoQueryRefund(outTradeNo string) (*allin.QueryRefundResult, error) {
	_, err := this.RefundQuery(RefundQueryRequest{OutTradeNo: outTradeNo})
	return nil, err
}

// DoPayMessage 支付通知
func (this WxPayPaymentService) DoPayMessage(queryStrings string) (*allin.PayMessageResult, error) {
	result := new(allin.PayMessageResult)
	result.Channel = core.WxPay
	rep, err := this.UnifiedOrderMessage(queryStrings)

	if rep != nil {
		result.PayAmount = rep.CashFee
		result.TotalAmount = rep.TotalFee
		result.OutTradeNo = rep.OutTradeNo
		result.TradeNo = rep.TransactionId
		result.TradeStatus = rep.TradeType
		result.BuyerId = rep.Openid
		// rep.NotifyTime
		result.NotifyTime = time.Now()
		result.Response = rep
	}
	return result, err
}

// DoRefundMessage 退款通知
func (this WxPayPaymentService) DoRefundMessage(queryStrings string) (*allin.RefundMessageResult, error) {
	_, err := this.PayRefundMessage(queryStrings)
	return nil, err
}
