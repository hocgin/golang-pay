package wxpay

import (
	"errors"
	"github.com/hocgin/golang-pay/core"
)

type WxPayPaymentService struct {
	ConfigStorage WxPayConfigStorage
}

func (this WxPayPaymentService) getUrl() string {
	if this.ConfigStorage.Ext.IsDev {
		return "https://api.mch.weixin.qq.com/sandboxnew"
	}
	return "https://api.mch.weixin.qq.com"
}

func (this WxPayPaymentService) RequestObject(request WxPayRequest, v interface{}) error {
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

func (this WxPayPaymentService) MessageObject(queryString string, v interface{}) error {
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
func (this WxPayPaymentService) UnifiedOrder(request WxPayRequest) (*UnifiedOrderResponse, error) {
	result := &UnifiedOrderResponse{}
	err := this.RequestObject(request, result)
	return result, err
}

// message
func (this WxPayPaymentService) PayRefund(queryString string) (*PayRefundMessage, error) {
	result := &PayRefundMessage{}
	err := this.MessageObject(queryString, result)
	return result, err
}
