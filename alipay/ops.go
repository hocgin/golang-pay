package alipay

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hocgin/golang-pay/core"
	"github.com/hocgin/golang-pay/core/net"
	"github.com/hocgin/golang-pay/core/utils"
	"net/url"
	"strings"
)

// request
type AliPayRequestImpl struct {
	Method              string `json:"method,omitempty"`
	AppId               string `json:"app_id,omitempty"`
	Format              string `json:"format,omitempty"`
	Charset             string `json:"charset,omitempty"`
	SignType            string `json:"sign_type,omitempty"`
	Sign                string `json:"sign,omitempty"`
	Timestamp           string `json:"timestamp,omitempty"`
	Version             string `json:"version,omitempty"`
	NotifyUrl           string `json:"notify_url,omitempty"`
	AppAuthToken        string `json:"app_auth_token,omitempty"`
	AliPayRequest       `json:"-"`
	core.PayRequestImpl `json:"-"`
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

func (this *AliPayRequestImpl) DoBuildUrl(request interface{}) (string, error) {
	service := this.PayService.(AliPayPaymentService)

	baseUrl := service.getUrl()
	bytes, err := json.Marshal(request)
	values := utils.JsonToMapValues(string(bytes))
	return net.GetUrlEncode(baseUrl, values), err
}

func (this *AliPayRequestImpl) DoBuildForm(request interface{}) (string, error) {
	service := this.PayService.(AliPayPaymentService)

	baseUrl := service.getUrl()
	bytes, err := json.Marshal(request)
	values := utils.JsonToMapValues(string(bytes))
	result := fmt.Sprintf(`<form method="POST" action="%s">`, baseUrl)
	for key, value := range values {
		result += fmt.Sprintf(`<input type="hidden" name='%s' value='%s'/>`, key, value)
	}
	result += `<input type="submit" value="立即支付" style="display:none"/></form><script>document.forms[0].submit();</script>`
	return result, err
}

// response
type AliPayResponseImpl struct {
	Code    string `json:"code,omitempty"`
	Msg     string `json:"msg,omitempty"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
	core.PayResponse
}

type AliPayResponse interface {
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

// message
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

func (this *AliPayMessageImpl) ToObject(v interface{}) error {
	queryString := this.MessageBody
	return utils.QueryParams(queryString, v)
}
