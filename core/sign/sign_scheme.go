package sign

type SignScheme interface {
	Name() string
	Sign(data string, privateKey string) string
	Verify(data string, publicKey string, sign string) bool
}
