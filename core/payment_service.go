package core

import "errors"

type PaymentServiceImpl struct {
}

func RequestObject(this interface{}, request interface{}, v interface{}) error {
	// ==================== [请求] ====================
	// 1. 构建请求
	if ref, isOk := request.(SetPayService); isOk {
		ref.SetPayService(this)
	}

	if ref, isOk := request.(AfterPropertiesSet); isOk {
		ref.AfterPropertiesSet()
	}

	// 2. 设置签名
	if ref, isOk := request.(FillSign); isOk {
		ref.FillSign(request)
	}

	// 3. 发起请求
	var body string
	var err error
	if ref, isOk := request.(DoRequest); isOk {
		body, err = ref.DoRequest(request)
		if err != nil {
			return err
		}
	}

	// ==================== [响应] ====================
	// 1. 响应数据
	if ref, isOk := v.(SetPayService); isOk {
		ref.SetPayService(this)
	}

	if ref, isOk := v.(SetBody); isOk {
		ref.SetBody(body)
	}

	// 2. 检查签名
	if ref, isOk := v.(CheckSign); isOk && !ref.CheckSign() {
		return errors.New("签名验证失败")
	}

	// 3. 转实体
	if ref, isOk := v.(ToObject); isOk {
		if err := ref.ToObject(v); err != nil {
			return err
		}
	}

	if response, isOk := v.(AfterPropertiesSet); isOk {
		response.AfterPropertiesSet()
	}
	return err
}

func MessageObject(this interface{}, queryString string, v interface{}) error {
	if ref, isOk := v.(SetPayService); isOk {
		ref.SetPayService(this)
	}

	if ref, isOk := v.(SetBody); isOk {
		ref.SetBody(queryString)
	}

	if ref, isOk := v.(CheckSign); isOk && !ref.CheckSign() {
		return errors.New("签名验证失败")
	}

	if ref, isOk := v.(ToObject); isOk {
		if err := ref.ToObject(v); err != nil {
			return err
		}
	}

	if ref, isOk := v.(AfterPropertiesSet); isOk {
		ref.AfterPropertiesSet()
	}
	return nil
}
