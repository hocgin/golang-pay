package allin

import (
	"github.com/hocgin/golang-pay/core"
	"time"
)

type PayWay int

const (
	// 支付宝
	AliPayApp PayWay = iota
	AliPayNative
	AliPayWeb
	AliPayQrCode
	// 微信支付
	WxPayJsApi
	WxPayApp
	WxPayNative
)

type TradeInfo struct {
	// [必选]
	Subject     string  // 标题
	TotalAmount float64 // 支付金额
	OutTradeNo  string  // 支付单号
	PayWay      PayWay  // 支付方式
	// [可选]
	NotifyUrl string // 通知地址
	// #AliPay
	QuitUrl string
	// #Openid
	Openid string
	Ext    map[string]string
}

type RefundInfo struct {
	RefundAmount float64      // 退款金额
	OutTradeNo   string       // 支付单号
	channel      core.Channel //
}

func GetServiceName(payWay PayWay) core.Channel {
	return ""
}

type PayResult struct {
	PayWay   PayWay
	Result   string
	Response interface{}
}
type QueryPayResult struct {
}
type QueryRefundResult struct {
}
type PayMessageResult struct {
	Channel     core.Channel // 来源
	OutTradeNo  string       // 交易单号
	TradeNo     string       // 平台交易单号
	NotifyTime  time.Time    // 通知时间
	BuyerId     string       // 买家标记
	PayAmount   float64      // 支付金额
	TotalAmount float64      // 交易总金额
	TradeStatus string       // 交易状态
	Response    interface{}
}
type RefundMessageResult struct {
}
