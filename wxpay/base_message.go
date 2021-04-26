package wxpay

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/antchfx/xpath"
	"github.com/hocgin/golang-pay/core"
	"strings"
)

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

func (this WxPayMessageImpl) ToObject(v interface{}) error {
	if err := xml.Unmarshal([]byte(this.MessageBody), &v); err != nil {
		return err
	}
	return nil
}

func (this WxPayMessageImpl) getXmlValue(path string) string {
	doc, err := this.getXmlDoc("xml")
	if err != nil {
		panic(err)
	}
	if textNode := doc.SelectElement(fmt.Sprintf("//%s//text()", path)); textNode != nil {
		return textNode.InnerText()
	}
	panic(errors.New("解析错误"))
}

func (this WxPayMessageImpl) getXmlDoc(rootName string) (*xmlquery.Node, error) {
	doc, err := xmlquery.Parse(strings.NewReader(this.MessageBody))
	if err != nil {
		return nil, err
	}
	root := xmlquery.FindOne(doc, fmt.Sprintf("//%s", rootName))
	return root, nil
}

// xml转map
func (this WxPayMessageImpl) toMaps() map[string]interface{} {
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
