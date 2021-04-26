package alipay

import (
	"fmt"
	"github.com/hocgin/golang-pay/core"
	"github.com/hocgin/golang-pay/core/utils"
	"net/url"
	"strings"
)

type AliPayMessageImpl struct {
	SignType string `json:"sign_type,omitempty"`
	Sign     string `json:"sign,omitempty"`
	core.PayMessage
}

type AliPayMessage interface {
}

func (this *AliPayMessageImpl) CheckSign() bool {
	service := this.PayService.(AliPayPaymentService)
	storage := service.ConfigStorage
	publicKey := storage.PublicKey

	queryString := this.MessageBody
	newQueryParams, _ := url.QueryUnescape(queryString)
	queryStrings := utils.QueryStrings(queryString)
	signType := queryStrings[SIGN_TYPE_FIELD_NAME].(string)
	sign := queryStrings[SIGN_FIELD_NAME].(string)
	waitSignString := strings.ReplaceAll(newQueryParams, fmt.Sprintf(`&%s=%s`, SIGN_FIELD_NAME, sign), "")
	waitSignString = strings.ReplaceAll(newQueryParams, fmt.Sprintf(`&%s=%s`, SIGN_TYPE_FIELD_NAME, signType), "")

	vars := utils.QueryStrings(waitSignString)
	keys := utils.KeysOrdered(utils.Keys(vars), false)
	unsignStr := utils.ConnectEncode(vars, keys, "&")
	signScheme := GetSignScheme(signType)

	return signScheme.Verify(unsignStr, publicKey, sign)
}

func (this AliPayMessageImpl) ToObject(v interface{}) error {
	queryString := this.MessageBody
	return utils.QueryParams(queryString, v)
}
