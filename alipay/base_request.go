package alipay

import (
	"encoding/json"
	"github.com/hocgin/golang-pay/core"
	"github.com/hocgin/golang-pay/core/net"
	"github.com/hocgin/golang-pay/core/utils"
)

type AliPayRequestImpl struct {
	Method        string `json:"method,omitempty"`
	AppId         string `json:"app_id,omitempty"`
	Format        string `json:"format,omitempty"`
	Charset       string `json:"charset,omitempty"`
	SignType      string `json:"sign_type,omitempty"`
	Sign          string `json:"sign,omitempty"`
	Timestamp     string `json:"timestamp,omitempty"`
	Version       string `json:"version,omitempty"`
	NotifyUrl     string `json:"notify_url,omitempty"`
	AppAuthToken  string `json:"app_auth_token,omitempty"`
	AliPayRequest `json:"-"`
	core.PayRequest
}

type AliPayRequest interface {
}

func (this *AliPayRequestImpl) AfterPropertiesSet() {
	service := this.PayService.(AliPayPaymentService)

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
func (this *AliPayRequestImpl) FillSign(request interface{}) {
	service := this.PayService.(AliPayPaymentService)
	storage := service.ConfigStorage

	signType := storage.Ext.SignType
	privateKey := storage.PrivateKey

	bytes, _ := json.Marshal(request)
	data := string(bytes)
	values := utils.JsonToMapValues(data)
	this.Sign = GetSign(values, signType, privateKey)
}

func (this *AliPayRequestImpl) DoRequest(request interface{}) (string, error) {
	service := this.PayService.(AliPayPaymentService)

	baseUrl := service.getUrl()
	bytes, _ := json.Marshal(request)
	values := utils.JsonToMapValues(string(bytes))

	url := net.GetUrlEncode(baseUrl, values)
	return net.GetString(url)
}
