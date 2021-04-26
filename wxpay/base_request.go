package wxpay

import (
	"encoding/xml"
	"github.com/hocgin/golang-pay/core"
	"github.com/hocgin/golang-pay/core/net"
	"github.com/hocgin/golang-pay/core/utils"
)

type WxPayRequestImpl struct {
	AppId           string `xml:"appid,omitempty"`
	MchId           string `xml:"mch_id,omitempty"`
	NonceStr        string `xml:"nonce_str,omitempty"`
	Sign            string `xml:"sign,omitempty"`
	SignType        string `xml:"sign_type"`
	RequestUrl      string `xml:"-"`
	core.PayRequest `xml:"-"`
	WxPayRequest
}

type WxPayRequest interface {
}

func (this *WxPayRequestImpl) FillSign(request interface{}) {
	service := this.PayService.(WxPayPaymentService)
	storage := service.ConfigStorage

	values := utils.XmlToMapValues(request)
	signValue := GetSign(values, storage.Ext.SignType, storage.Key)
	this.Sign = signValue
}

func (this *WxPayRequestImpl) DoRequest(request interface{}) (string, error) {
	service := this.PayService.(WxPayPaymentService)
	bytes, err := xml.Marshal(request)
	if err != nil {
		return "", err
	}
	requestBody := string(bytes)
	return net.PostString(service.getUrl()+this.RequestUrl, requestBody)
}

func (this *WxPayRequestImpl) AfterPropertiesSet() {
	service := this.PayService.(WxPayPaymentService)
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
