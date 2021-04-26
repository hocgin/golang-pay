package wxpay

import (
	"encoding/xml"
	"github.com/hocgin/golang-pay/core"
)

type WxPayResponseImpl struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	Appid      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
	Sign       string `xml:"sign"`
	core.PayResponse
}

type WxPayResponse interface {
}

func (this *WxPayResponseImpl) ToObject(v interface{}) error {
	if err := xml.Unmarshal([]byte(this.ResponseBody), &v); err != nil {
		return err
	}
	return nil
}

func (this *WxPayResponseImpl) IsCheckSign() bool {
	return true
}
func (this *WxPayResponseImpl) CheckSign() bool {
	service := this.PayService.(WxPayPaymentService)
	storage := service.ConfigStorage
	key := storage.Key
	signType := storage.Ext.SignType

	body := this.ResponseBody
	bodyMap := make(map[string]interface{})
	_ = xml.Unmarshal([]byte(body), (*StringMap)(&bodyMap))

	verifySignValue := bodyMap[SIGN_FIELD_NAME].(string)
	bodyMap[SIGN_FIELD_NAME] = ""
	return signType.Verify(GetSignValue(bodyMap, key), key, verifySignValue)
}
