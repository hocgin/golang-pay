package wxpay

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/antchfx/xpath"
	"github.com/hocgin/golang-pay/core"
	"github.com/hocgin/golang-pay/core/net"
	"github.com/hocgin/golang-pay/core/utils"
	"strings"
)

// request
type WxPayRequestImpl struct {
	AppId               string `xml:"appid,omitempty"`
	MchId               string `xml:"mch_id,omitempty"`
	NonceStr            string `xml:"nonce_str,omitempty"`
	Sign                string `xml:"sign,omitempty"`
	SignType            string `xml:"sign_type"`
	RequestUrl          string `xml:"-"`
	core.PayRequestImpl `xml:"-"`
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

// response
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

// message
type WxPayMessageImpl struct {
	core.PayMessage
}

type WxPayMessage interface {
}

func (this *WxPayMessageImpl) CheckSign() bool {
	service := this.PayService.(WxPayPaymentService)
	storage := service.ConfigStorage
	key := storage.Key
	signType := storage.Ext.SignType

	body := this.MessageBody
	bodyMap := make(map[string]interface{})
	_ = xml.Unmarshal([]byte(body), (*StringMap)(&bodyMap))
	verifySignValue := bodyMap[SIGN_FIELD_NAME].(string)
	bodyMap[SIGN_FIELD_NAME] = ""
	return signType.Verify(GetSignValue(bodyMap, key), key, verifySignValue)
}

func (this *WxPayMessageImpl) ToObject(v interface{}) error {
	if err := xml.Unmarshal([]byte(this.MessageBody), &v); err != nil {
		return err
	}
	return nil
}

func (this *WxPayMessageImpl) getXmlValue(path string) string {
	doc, err := this.getXmlDoc("xml")
	if err != nil {
		panic(err)
	}
	if textNode := doc.SelectElement(fmt.Sprintf("//%s//text()", path)); textNode != nil {
		return textNode.InnerText()
	}
	panic(errors.New("解析错误"))
}

func (this *WxPayMessageImpl) getXmlDoc(rootName string) (*xmlquery.Node, error) {
	doc, err := xmlquery.Parse(strings.NewReader(this.MessageBody))
	if err != nil {
		return nil, err
	}
	root := xmlquery.FindOne(doc, fmt.Sprintf("//%s", rootName))
	return root, nil
}

// xml转map
func (this *WxPayMessageImpl) toMaps() map[string]interface{} {
	result := make(map[string]interface{})
	root, err := this.getXmlDoc("xml")
	if err != nil {
		return result
	}
	expr, err := xpath.Compile("//xml/*")
	if err != nil {
		return result

	}
	iter := expr.Select(xmlquery.CreateXPathNavigator(root))
	for iter.MoveNext() {
		current := iter.Current()
		result[current.LocalName()] = current.Value()
	}
	return result
}
