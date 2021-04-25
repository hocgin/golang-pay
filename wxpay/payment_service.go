package wxpay

import (
    "encoding/xml"
    "errors"
    "github.com/hocgin/golang-pay/core/net"
    "github.com/hocgin/golang-pay/core/ops"
    "github.com/hocgin/golang-pay/core/utils"
)

type WxPayPaymentService struct {
    ConfigStorage WxPayConfigStorage
}

func (this WxPayPaymentService) getUrl() string {
    if this.ConfigStorage.Ext.IsDev {
        return "https://api.mch.weixin.qq.com/sandboxnew"
    }
    return "https://api.mch.weixin.qq.com"
}

func (this WxPayPaymentService) RequestObject(request WxPayRequest, v interface{}) error {
    storage := this.ConfigStorage
    key := storage.Key
    signType := storage.Ext.SignType
    baseUrl := this.getUrl()

    // 1. 构建参数
    request.DefaultConfig(this)
    request.RequestBefore()

    // 2. 签名
    values := utils.XmlToMapValues(request)
    signValue := GetSign(values, signType, key)
    request.SetSign(signValue)

    // 3. 构建URL
    url := (baseUrl + request.GetUrl())

    // 4. 发起请求
    bytes, _ := xml.Marshal(request)
    data := string(bytes)
    body, err := net.PostString(url, data)
    if err != nil {
        return err
    }

    // ==== [响应] ====
    // 1. 解析响应数据
    bodyMap := make(map[string]interface{})
    _ = xml.Unmarshal([]byte(body), (*StringMap)(&bodyMap))

    // 2. 检查签名
    verifySignValue := bodyMap[SIGN_FIELD_NAME].(string)
    bodyMap[SIGN_FIELD_NAME] = ""
    isOk := signType.Verify(GetSignValue(bodyMap, key), key, verifySignValue)
    if !isOk {
        return errors.New("验证签名失败")
    }

    // 3. 转实体
    err = xml.Unmarshal([]byte(body), &v)
    if response, isOk := v.(ops.SetResponseBody); isOk {
        response.SetResponseBody(body)
    }
    if response, isOk := v.(ops.AfterPropertiesSet); isOk {
        response.AfterPropertiesSet()
    }
    return err
}

func (this WxPayPaymentService) UnifiedOrder(request WxPayRequest) (*UnifiedOrderResponse, error) {
    result := &UnifiedOrderResponse{}
    err := this.RequestObject(request, result)
    return result, err
}
