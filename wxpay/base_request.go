package wxpay

import "github.com/hocgin/golang-pay/core/ops"

type WxPayRequestImpl struct {
	AppId    string `xml:"appid,omitempty"`
	MchId    string `xml:"mch_id,omitempty"`
	NonceStr string `xml:"nonce_str,omitempty"`
	Sign     string `xml:"sign,omitempty"`
	SignType string `xml:"sign_type,omitempty"`
}

type WxPayRequest interface {
	DefaultConfig(service WxPayPaymentService)
	GetUrl() string
	SetSign(sign string)
	ops.PayRequest
}

func (this *WxPayRequestImpl) SetSign(sign string) {
	this.Sign = sign
}

func (this *WxPayRequestImpl) RequestBefore() {
	// no thing
}

func (this *WxPayRequestImpl) DefaultConfig(service WxPayPaymentService) {
	storage := service.ConfigStorage
	if this.AppId == "" {
		this.AppId = storage.AppId
	}
	if this.MchId == "" {
		this.MchId = storage.MchId
	}
	if this.NonceStr == "" {
		this.NonceStr = GetNowTime()
	}
	if this.SignType == "" {
		this.SignType = GetSignTypeName(storage.Ext.SignType)
	}
}
