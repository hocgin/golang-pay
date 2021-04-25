package core

type ExtOption struct {
}

type DialOption interface {
	apply(*ExtOption)
}

type FunOption struct {
	f func(*ExtOption)
}

func (fdo *FunOption) apply(do *ExtOption) {
	fdo.f(do)
}

func NewFunctionOption(f func(o *ExtOption)) *FunOption {
	return &FunOption{f: f}
}
