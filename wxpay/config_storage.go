package wxpay

import (
	core "github.com/hocgin/golang-pay/core"
	"github.com/hocgin/golang-pay/core/sign"
)

type WxPayConfigStorage struct {
	Name  string
	AppId string
	MchId string
	Key   string

	Ext wxPayConfigStorageExt
	core.ConfigStorage
}

type wxPayConfigStorageExt struct {
	SignType     sign.SignScheme // default value: MD5
	CertFileText string          // default value: ""
	IsDev        bool            // default value: false
}

var DefaultExt = wxPayConfigStorageExt{
	SignType:     &sign.MD5{},
	IsDev:        false,
	CertFileText: "",
}

func (this *WxPayConfigStorage) CreateService() *WxPayPaymentService {
	return &WxPayPaymentService{ConfigStorage: *this}
}

type Option func(ext *wxPayConfigStorageExt)

func CreateConfigStorage(appId string, mchId string, key string, options ...func(*wxPayConfigStorageExt)) *WxPayConfigStorage {
	result := &WxPayConfigStorage{
		AppId: appId,
		MchId: mchId,
		Key:   key,
		Ext:   DefaultExt,
	}
	result.Name = "wxpay"
	for _, option := range options {
		option(&result.Ext)
	}
	return result
}

func IsDev(isDev bool) Option {
	return func(o *wxPayConfigStorageExt) {
		o.IsDev = isDev
	}
}
func CertFileText(certFileText string) Option {
	return func(o *wxPayConfigStorageExt) {
		o.CertFileText = certFileText
	}
}
