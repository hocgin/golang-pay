package ops

type PayResponse struct {
}

type SetResponseBody interface {
    SetResponseBody(body string)
}

type AfterPropertiesSet interface {
    AfterPropertiesSet()
}

type IsCheckSign interface {
    IsCheckSign() bool
}
