package alipay

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hocgin/golang-pay/core"
	"strings"
)

type AliPayResponseImpl struct {
	Code    string `json:"code,omitempty"`
	Msg     string `json:"msg,omitempty"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
	core.PayResponse
}

type AliPayResponse interface {
}

func (this *AliPayResponseImpl) IsCheckSign() bool {
	return true
}

func (this *AliPayResponseImpl) CheckSign() bool {
	service := this.PayService.(AliPayPaymentService)
	storage := service.ConfigStorage
	signType := storage.Ext.SignType
	publicKey := storage.PublicKey

	body := this.ResponseBody
	bodyMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(body), &bodyMap)
	if err != nil {
		return false
	}

	substr := "response"
	beginIndex := strings.Index(body, substr) + len(substr) + 2
	endIndex := strings.Index(body, fmt.Sprintf(`,"%s"`, SIGN_FIELD_NAME))
	if endIndex < 0 {
		return false
	}
	response := string([]byte(body)[beginIndex:endIndex])
	sign := bodyMap[SIGN_FIELD_NAME].(string)
	return signType.Verify(response, publicKey, sign)
}

func (this *AliPayResponseImpl) ToObject(v interface{}) error {
	body := this.ResponseBody
	bodyMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(body), &bodyMap)
	if err != nil {
		return err
	}

	substr := "response"
	beginIndex := strings.Index(body, substr) + len(substr) + 2
	endIndex := strings.Index(body, fmt.Sprintf(`,"%s"`, SIGN_FIELD_NAME))
	if endIndex < 0 {
		return errors.New("响应解析错误")
	}
	response := string([]byte(body)[beginIndex:endIndex])
	if err := json.Unmarshal([]byte(response), &v); err != nil {
		return err
	}
	return nil
}
