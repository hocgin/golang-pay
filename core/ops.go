package core

type AfterPropertiesSet interface {
	AfterPropertiesSet()
}

type SetBody interface {
	SetBody(body string)
}

type FillSign interface {
	FillSign(interface{})
}

type BeforePropertiesSet interface {
	BeforePropertiesSet()
}

type CheckSign interface {
	CheckSign() bool
}

type SetPayService interface {
	SetPayService(interface{})
}

type ToObject interface {
	ToObject(v interface{}) error
}

type DoRequest interface {
	DoRequest(interface{}) (string, error)
}

type DoBuildUrl interface {
	DoBuildUrl(interface{}) (string, error)
}

type DoBuildForm interface {
	DoBuildForm(interface{}) (string, error)
}

type Download interface {
}

// request
type PayRequestImpl struct {
	PayService interface{}
}

func (this *PayRequestImpl) SetPayService(payService interface{}) {
	this.PayService = payService
}

// response
type PayResponse struct {
	ResponseBody string
	PayService   interface{}
}

func (this *PayResponse) SetPayService(payService interface{}) {
	this.PayService = payService
}

func (this *PayResponse) SetBody(body string) {
	this.ResponseBody = body
}

// message
type PayMessage struct {
	MessageBody string
	PayService  interface{}
}

func (this *PayMessage) SetPayService(payService interface{}) {
	this.PayService = payService
}

func (this *PayMessage) SetBody(body string) {
	this.MessageBody = body
}
