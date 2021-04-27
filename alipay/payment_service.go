package alipay

import (
	"errors"
	"github.com/hocgin/golang-pay/core"
)

type AliPayPaymentService struct {
	ConfigStorage AliPayConfigStorage
	core.PaymentServiceImpl
}

func (service *AliPayPaymentService) getUrl() string {
	if service.ConfigStorage.Ext.IsDev {
		return "https://openapi.alipaydev.com/gateway.do"
	}
	return "https://openapi.alipay.com/gateway.do"
}

func BuildSdkRequest(this interface{}, request interface{}) (string, error) {
	// ==================== [请求] ====================
	// 1. 构建请求
	if ref, isOk := request.(core.SetPayService); isOk {
		ref.SetPayService(this)
	}

	if ref, isOk := request.(core.AfterPropertiesSet); isOk {
		ref.AfterPropertiesSet()
	}

	// 2. 设置签名
	if ref, isOk := request.(core.FillSign); isOk {
		ref.FillSign(request)
	}

	if ref, isOk := request.(core.DoBuildUrl); isOk {
		body, err := ref.DoBuildUrl(request)
		return body, err
	}
	return "", errors.New("构建请求失败")
}

func BuildFormRequest(this interface{}, request interface{}) (string, error) {
	// ==================== [请求] ====================
	// 1. 构建请求
	if ref, isOk := request.(core.SetPayService); isOk {
		ref.SetPayService(this)
	}

	if ref, isOk := request.(core.AfterPropertiesSet); isOk {
		ref.AfterPropertiesSet()
	}

	// 2. 设置签名
	if ref, isOk := request.(core.FillSign); isOk {
		ref.FillSign(request)
	}

	if ref, isOk := request.(core.DoBuildForm); isOk {
		body, err := ref.DoBuildForm(request)
		return body, err
	}
	return "", errors.New("构建请求失败")
}

// request
func (this AliPayPaymentService) TradePay(request AliPayRequest) (*TradePayResponse, error) {
	result := &TradePayResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}
func (this AliPayPaymentService) TradeCreate(request AliPayRequest) (*TradeCreateResponse, error) {
	result := &TradeCreateResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}
func (this AliPayPaymentService) TradePreCreate(request AliPayRequest) (*TradePreCreateResponse, error) {
	result := &TradePreCreateResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}
func (this AliPayPaymentService) TradeQuery(request AliPayRequest) (*TradeQueryResponse, error) {
	result := &TradeQueryResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}
func (this AliPayPaymentService) TradeWapPay(request AliPayRequest) (string, error) {
	return BuildFormRequest(this, request)
}
func (this AliPayPaymentService) TradeRefund(request AliPayRequest) (*TradeRefundResponse, error) {
	result := &TradeRefundResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}
func (this AliPayPaymentService) TradePageRefund(request AliPayRequest) (*TradePageRefundResponse, error) {
	result := &TradePageRefundResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}
func (this AliPayPaymentService) TradePagePay(request AliPayRequest) (string, error) {
	return BuildFormRequest(this, request)
}
func (this AliPayPaymentService) TradeOrderSettle(request AliPayRequest) (*TradeOrderSettleResponse, error) {
	result := &TradeOrderSettleResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}
func (this AliPayPaymentService) TradeOrderInfoSync(request AliPayRequest) (*TradeOrderInfoSyncResponse, error) {
	result := &TradeOrderInfoSyncResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}
func (this AliPayPaymentService) TradeFastpayRefundQuery(request AliPayRequest) (*TradeFastpayRefundQueryResponse, error) {
	result := &TradeFastpayRefundQueryResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}
func (this AliPayPaymentService) TradeClose(request AliPayRequest) (*TradeCloseResponse, error) {
	result := &TradeCloseResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}
func (this AliPayPaymentService) TradeCancel(request AliPayRequest) (*TradeCancelResponse, error) {
	result := &TradeCancelResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}
func (this AliPayPaymentService) TradeAppPay(request AliPayRequest) (string, error) {
	return BuildSdkRequest(this, request)
}
func (this AliPayPaymentService) TradeAdvanceConsult(request AliPayRequest) (*TradeAdvanceConsultResponse, error) {
	result := &TradeAdvanceConsultResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}

// message
func (this AliPayPaymentService) TradeStatusSyncMessage(queryString string) (*TradeStatusSyncMessage, error) {
	result := &TradeStatusSyncMessage{}
	err := core.MessageObject(this, queryString, result)
	return result, err
}
