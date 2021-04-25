package wxpay

import "github.com/hocgin/golang-pay/core/ops"

type WxPayResponseImpl struct {
    ReturnCode   string `xml:"return_code"`
    ReturnMsg    string `xml:"return_msg"`
    Appid        string `xml:"appid"`
    MchId        string `xml:"mch_id"`
    NonceStr     string `xml:"nonce_str"`
    ErrCode      string `xml:"err_code"`
    ErrCodeDes   string `xml:"err_code_des"`
    Sign         string `xml:"sign"`
    ResponseBody string
    ops.PayResponse
}

type WxPayResponse interface {
}

func (this *WxPayResponseImpl) SetResponseBody(body string) {
    this.ResponseBody = body
}
