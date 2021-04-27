package wxpay

import (
	"errors"
	"fmt"
	"github.com/hocgin/golang-pay/core/sign"
	"github.com/hocgin/golang-pay/core/utils"
	"strconv"
	"strings"
	"time"
)

func GetSignTypeName(scheme sign.SignScheme) string {
	switch scheme.Name() {
	case "MD5":
		return "MD5"
	case "HMAC-SHA256":
		return "HMAC-SHA256"
	}
	panic(errors.New("签名算法没有找到"))
}

func GetSignScheme(signType string) sign.SignScheme {
	switch signType {
	case "MD5":
		return &sign.MD5{}
	case "HMAC-SHA256":
		return &sign.MD5{}
	}
	panic(errors.New("签名算法没有找到"))
}

func GetNowTime() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func GetSignValue(values map[string]interface{}, key string) string {
	vars := utils.MapFilterNullOrEmptry(values)
	keys := utils.KeysOrdered(utils.Keys(vars), false)
	unsignStr := utils.ConnectEncode(vars, keys, "&")
	return fmt.Sprintf(`%s&key=%s`, unsignStr, key)
}

func GetSign(values map[string]interface{}, signScheme sign.SignScheme, key string) string {
	signString := GetSignValue(values, key)
	signValue := signScheme.Sign(signString, key)
	return strings.ToUpper(signValue)
}
