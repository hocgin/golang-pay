package alipay

import (
	"errors"
	"github.com/hocgin/golang-pay/core/sign"
	"github.com/hocgin/golang-pay/core/utils"
	"time"
)

func GetSignTypeName(scheme sign.SignScheme) string {
	switch scheme.Name() {
	case "RSAWithSHA128":
		return "RSA"
	case "RSAWithSHA256":
		return "RSA2"
	}
	panic(errors.New("签名算法没有找到"))
}

func GetSignScheme(signType string) sign.SignScheme {
	switch signType {
	case "RSA2":
		return &sign.RSAWithSHA256{}
	case "RSA":
		return &sign.RSAWithSHA128{}
	}
	panic(errors.New("签名算法没有找到"))
}

func GetNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetSign(values map[string]interface{}, signScheme sign.SignScheme, privateKey string) string {
	vars := utils.MapFilterNullOrEmptry(values)
	keys := utils.KeysOrdered(utils.Keys(vars), false)
	unsignStr := utils.ConnectEncode(vars, keys, "&")
	signValue := signScheme.Sign(unsignStr, privateKey)
	return signValue
}
