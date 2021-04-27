package allin

import (
	"github.com/hocgin/golang-pay/alipay"
	"github.com/hocgin/golang-pay/core"
	"github.com/hocgin/golang-pay/wxpay"
)

type AllInPayServiceImpl struct {
	services map[core.Channel]AllInPayService
}

func (this AllInPayServiceImpl) Configs(services ...AllInPayService) {
	for _, service := range services {
		switch service.(type) {
		case alipay.AliPayPaymentService:
			this.services[core.AliPay] = service
		case wxpay.WxPayPaymentService:
			this.services[core.WxPay] = service
		}
	}
}

// DoPay 去支付
func (this AllInPayServiceImpl) DoPay(trade TradeInfo) (*PayResult, error) {
	return this.services[GetServiceName(trade.PayWay)].DoPay(trade)
}

// DoRefund 去退款
func (this AllInPayServiceImpl) DoRefund(refund RefundInfo) error {
	return this.services[refund.channel].DoRefund(refund)
}

// DoQueryPay 查询支付
func (this AllInPayServiceImpl) DoQueryPay(channel core.Channel, outTradeNo string) (*QueryPayResult, error) {
	return this.services[channel].DoQueryPay(outTradeNo)
}

// DoQueryRefund 查询退款
func (this AllInPayServiceImpl) DoQueryRefund(channel core.Channel, outTradeNo string) (*QueryRefundResult, error) {
	return this.services[channel].DoQueryRefund(outTradeNo)
}

// DoPayMessage 支付通知
func (this AllInPayServiceImpl) DoPayMessage(channel core.Channel, queryStrings string) (*PayMessageResult, error) {
	return this.services[channel].DoPayMessage(queryStrings)
}

// DoRefundMessage 退款通知
func (this AllInPayServiceImpl) DoRefundMessage(channel core.Channel, queryStrings string) (*RefundMessageResult, error) {
	return this.services[channel].DoRefundMessage(queryStrings)
}
