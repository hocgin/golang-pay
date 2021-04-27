package wxpay

import (
	"github.com/hocgin/golang-pay/core"
)

type WxPayPaymentService struct {
	ConfigStorage WxPayConfigStorage
	core.PaymentServiceImpl
}

func (this WxPayPaymentService) getUrl() string {
	if this.ConfigStorage.Ext.IsDev {
		return "https://api.mch.weixin.qq.com/sandboxnew"
	}
	return "https://api.mch.weixin.qq.com"
}

// request
func (this WxPayPaymentService) UnifiedOrder(request WxPayRequest) (*UnifiedOrderResponse, error) {
	result := &UnifiedOrderResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}
func (this WxPayPaymentService) RefundQuery(request WxPayRequest) (*RefundQueryResponse, error) {
	result := &RefundQueryResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}
func (this WxPayPaymentService) PayitilReport(request WxPayRequest) (*PayitilReportResponse, error) {
	result := &PayitilReportResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}
func (this WxPayPaymentService) PayRefund(request WxPayRequest) (*PayRefundResponse, error) {
	result := &PayRefundResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}
func (this WxPayPaymentService) OrderQuery(request WxPayRequest) (*OrderQueryResponse, error) {
	result := &OrderQueryResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}
func (this WxPayPaymentService) DownloadFundFlow(request WxPayRequest) (*DownloadFundFlowResponse, error) {
	result := &DownloadFundFlowResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}
func (this WxPayPaymentService) DownloadBill(request WxPayRequest) (*DownloadBillResponse, error) {
	result := &DownloadBillResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}
func (this WxPayPaymentService) CloseOrder(request WxPayRequest) (*CloseOrderResponse, error) {
	result := &CloseOrderResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}
func (this WxPayPaymentService) BatchQueryComment(request WxPayRequest) (*BatchQueryCommentResponse, error) {
	result := &BatchQueryCommentResponse{}
	err := core.RequestObject(this, request, result)
	return result, err
}

// message
func (this WxPayPaymentService) PayRefundMessage(queryString string) (*PayRefundMessage, error) {
	result := &PayRefundMessage{}
	err := core.MessageObject(this, queryString, result)
	return result, err
}
func (this WxPayPaymentService) UnifiedOrderMessage(queryString string) (*UnifiedOrderMessage, error) {
	result := &UnifiedOrderMessage{}
	err := core.MessageObject(this, queryString, result)
	return result, err
}
