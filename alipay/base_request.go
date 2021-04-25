package alipay

import "github.com/hocgin/golang-pay/core/ops"

type AliPayRequestImpl struct {
	Method       string `json:"method,omitempty"`
	AppId        string `json:"app_id,omitempty"`
	Format       string `json:"format,omitempty"`
	Charset      string `json:"charset,omitempty"`
	SignType     string `json:"sign_type,omitempty"`
	Sign         string `json:"sign,omitempty"`
	Timestamp    string `json:"timestamp,omitempty"`
	Version      string `json:"version,omitempty"`
	NotifyUrl    string `json:"notify_url,omitempty"`
	AppAuthToken string `json:"app_auth_token,omitempty"`
}

type AliPayRequest interface {
	DefaultConfig(service AliPayPaymentService)
	ops.PayRequest
	GetMethod() string
}

func (this *AliPayRequestImpl) GetMethod() string {
	return this.Method
}

func (this *AliPayRequestImpl) RequestBefore() {
	// no thing
}

func (this *AliPayRequestImpl) DefaultConfig(service AliPayPaymentService) {
	storage := service.ConfigStorage
	if this.AppId == "" {
		this.AppId = storage.AppId
	}
	if this.SignType == "" {
		this.SignType = GetSignTypeName(storage.Ext.SignType)
	}
	if this.Version == "" {
		this.Version = storage.Ext.Version
	}
	if this.Format == "" {
		this.Format = storage.Ext.Format
	}
	if this.Charset == "" {
		this.Charset = storage.Ext.Charset
	}
	if this.Timestamp == "" {
		this.Timestamp = GetNowTime()
	}
}
