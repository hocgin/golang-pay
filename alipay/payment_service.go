package alipay

import (
	"encoding/json"
	"errors"
	"github.com/hocgin/golang-pay/core/net"
	"github.com/hocgin/golang-pay/core/utils"
	"net/url"
	"strings"
)

type AliPayPaymentService struct {
	ConfigStorage AliPayConfigStorage
}

func (service *AliPayPaymentService) getUrl() string {
	if service.ConfigStorage.Ext.IsDev {
		return "https://openapi.alipaydev.com/gateway.do"
	}
	return "https://openapi.alipay.com/gateway.do"
}
func (this AliPayPaymentService) RequestObject(request AliPayRequest, v interface{}) error {
	storage := this.ConfigStorage
	signType := storage.Ext.SignType
	privateKey := storage.PrivateKey
	publicKey := storage.PublicKey
	baseUrl := this.getUrl()

	// 1. 构建参数
	request.DefaultConfig(this)
	request.RequestBefore()
	bytes, _ := json.Marshal(request)
	data := string(bytes)

	// 2. 签名
	values := utils.JsonToMapValues(data)
	signValue := GetSign(values, signType, privateKey)
	values[SIGN_FIELD_NAME] = signValue

	// 3. 构建URL
	url := net.GetUrlEncode(baseUrl, values)

	// 4. 发起请求
	body, err := net.GetString(url)
	if err != nil {
		return err
	}

	// ==== [响应] ====
	// 1. 解析响应数据
	bodyMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(body), &bodyMap)
	if err != nil {
		return err
	}

	substr := "response"
	beginIndex := strings.Index(body, substr) + len(substr) + 2
	endIndex := strings.Index(body, ",\""+SIGN_FIELD_NAME+" \"")
	if endIndex < 0 {
		return errors.New("响应解析错误")
	}
	response := string([]byte(body)[beginIndex:endIndex])
	sign := bodyMap[SIGN_FIELD_NAME].(string)

	// 2. 检查签名
	isOk := signType.Verify(response, publicKey, sign)
	if !isOk {
		return errors.New("验证签名失败")
	}
	_ = json.Unmarshal([]byte(response), v)
	return err
}

func (this AliPayPaymentService) MessageObject(queryParams string, v interface{}) error {
	publicKey := this.ConfigStorage.PublicKey

	newQueryParams, _ := url.QueryUnescape(queryParams)
	queryStrings := utils.QueryStrings(queryParams)
	signType := queryStrings[SIGN_TYPE_FIELD_NAME].(string)
	sign := queryStrings[SIGN_FIELD_NAME].(string)
	waitSignString := strings.ReplaceAll(newQueryParams, "&"+SIGN_FIELD_NAME+"="+sign, "")
	waitSignString = strings.ReplaceAll(waitSignString, "&"+SIGN_TYPE_FIELD_NAME+"="+signType, "")

	vars := utils.QueryStrings(waitSignString)
	keys := utils.KeysOrdered(utils.Keys(vars), false)
	unsignStr := utils.ConnectEncode(vars, keys, "&")
	signScheme := GetSignScheme(signType)

	if !signScheme.Verify(unsignStr, publicKey, sign) {
		return errors.New("校验失败")
	}
	err := utils.QueryParams(queryParams, v)
	return err
}

// request
func (this AliPayPaymentService) TradePay(request AliPayRequest) (*TradePayResponse, error) {
	result := &TradePayResponse{}
	err := this.RequestObject(request, result)
	return result, err
}
func (this AliPayPaymentService) TradeCreate(request AliPayRequest) (*TradeCreateResponse, error) {
	result := &TradeCreateResponse{}
	err := this.RequestObject(request, result)
	return result, err
}
func (this AliPayPaymentService) TradePreCreate(request AliPayRequest) (*TradePreCreateResponse, error) {
	result := &TradePreCreateResponse{}
	err := this.RequestObject(request, result)
	return result, err
}
func (this AliPayPaymentService) TradeQuery(request AliPayRequest) (*TradeQueryResponse, error) {
	result := &TradeQueryResponse{}
	err := this.RequestObject(request, result)
	return result, err
}

// message
func (this AliPayPaymentService) TradeStatusSyncMessage(queryParams string) (*TradeStatusSyncMessage, error) {
	result := &TradeStatusSyncMessage{}
	err := this.MessageObject(queryParams, result)
	return result, err
}
