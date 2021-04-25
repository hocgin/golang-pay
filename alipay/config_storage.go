package alipay

import (
	"github.com/hocgin/golang-pay/core"
	"github.com/hocgin/golang-pay/core/sign"
)

type AliPayConfigStorage struct {
	Name       string
	PublicKey  string
	PrivateKey string
	AppId      string

	Ext aliPayConfigStorageExt
	core.ConfigStorage
}

type aliPayConfigStorageExt struct {
	SignType sign.SignScheme // default value: RSA2
	IsDev    bool            // default value: true
	Charset  string          // default value: utf-8
	Format   string          // default value: json
	Version  string          // default value: 1.0
}

func CreateConfigStorage(appId string, publicKey string, privateKey string, options ...func(*aliPayConfigStorageExt)) *AliPayConfigStorage {
	result := &AliPayConfigStorage{
		Name:       "alipay",
		AppId:      appId,
		PublicKey:  publicKey,
		PrivateKey: privateKey,
		Ext:        DefaultExt,
	}
	for _, option := range options {
		option(&result.Ext)
	}
	return result
}

func (this *AliPayConfigStorage) CreateService() *AliPayPaymentService {
	return &AliPayPaymentService{ConfigStorage: *this}
}

var DefaultExt = aliPayConfigStorageExt{
	SignType: &sign.RSAWithSHA256{},
	IsDev:    false,
	Charset:  "utf-8",
	Format:   "json",
	Version:  "1.0",
}

type Option func(*aliPayConfigStorageExt)

func SignType(signType sign.SignScheme) Option {
	return func(o *aliPayConfigStorageExt) {
		o.SignType = signType
	}
}
func Version(version string) Option {
	return func(o *aliPayConfigStorageExt) {
		o.Version = version
	}
}
func Format(format string) Option {
	return func(o *aliPayConfigStorageExt) {
		o.Format = format
	}
}
func IsDev(isDev bool) Option {
	return func(o *aliPayConfigStorageExt) {
		o.IsDev = isDev
	}
}
func Charset(charset string) Option {
	return func(o *aliPayConfigStorageExt) {
		o.Charset = charset
	}
}
