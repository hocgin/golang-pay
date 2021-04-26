package alipay

import (
	"errors"
	"github.com/hocgin/golang-pay/core"
)

type AliPayPaymentService struct {
	ConfigStorage AliPayConfigStorage
}

func (service *AliPayPaymentService) getUrl() string {
	if service.ConfigStorage.Ext.IsDev {
		return "https://openapi.alipaydev.com/gateway.do"
	}
	return "https://openapi.alipay.com/gateway.do"
}
func (this AliPayPaymentService) RequestObject(request AliPayRequest, v interface{}) error {
	// ==================== [请求] ====================
	// 1. 构建参数
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

	// 3. 发起请求
	var body string
	var err error
	if ref, isOk := request.(core.DoRequest); isOk {
		body, err = ref.DoRequest(request)
		if err != nil {
			return err
		}
	}

	// ==================== [响应] ====================
	// 1. 响应数据
	if ref, isOk := v.(core.SetPayService); isOk {
		ref.SetPayService(this)
	}

	if ref, isOk := v.(core.SetBody); isOk {
		ref.SetBody(body)
	}

	// 2. 检查签名
	if ref, isOk := v.(core.IsCheckSign); isOk && ref.IsCheckSign() {
		if ref, isOk := v.(core.CheckSign); isOk && !ref.CheckSign() {
			return errors.New("签名验证失败")
		}
	}

	// 3. 转实体
	if ref, isOk := v.(core.ToObject); isOk {
		if err := ref.ToObject(v); err != nil {
			return err
		}
	}
	if response, isOk := v.(core.AfterPropertiesSet); isOk {
		response.AfterPropertiesSet()
	}
	return err
}

func (this AliPayPaymentService) MessageObject(queryString string, v interface{}) error {
	if ref, isOk := v.(core.SetPayService); isOk {
		ref.SetPayService(this)
	}

	if ref, isOk := v.(core.SetBody); isOk {
		ref.SetBody(queryString)
	}

	if ref, isOk := v.(core.IsCheckSign); isOk && ref.IsCheckSign() {
		if ref, isOk := v.(core.CheckSign); isOk && !ref.CheckSign() {
			return errors.New("签名验证失败")
		}
	}

	if ref, isOk := v.(core.ToObject); isOk {
		if err := ref.ToObject(v); err != nil {
			return err
		}
	}

	if ref, isOk := v.(core.AfterPropertiesSet); isOk {
		ref.AfterPropertiesSet()
	}
	return nil
}

// request
func (this AliPayPaymentService) TradePay(request AliPayRequest) (*TradePayResponse, error) {
	result := &TradePayResponse{}
	err := this.RequestObject(request, result)
	return result, err
}
func (this AliPayPaymentService) TradeCreate(request AliPayRequest) (*TradeCreateResponse, error) {
	result := &TradeCreateResponse{}
	err := this.RequestObject(request, result)
	return result, err
}
func (this AliPayPaymentService) TradePreCreate(request AliPayRequest) (*TradePreCreateResponse, error) {
	result := &TradePreCreateResponse{}
	err := this.RequestObject(request, result)
	return result, err
}
func (this AliPayPaymentService) TradeQuery(request AliPayRequest) (*TradeQueryResponse, error) {
	result := &TradeQueryResponse{}
	err := this.RequestObject(request, result)
	return result, err
}

// message
func (this AliPayPaymentService) TradeStatusSyncMessage(queryString string) (*TradeStatusSyncMessage, error) {
	result := &TradeStatusSyncMessage{}
	err := this.MessageObject(queryString, result)
	return result, err
}
