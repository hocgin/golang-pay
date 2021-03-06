package sign

import (
	"testing"
)

func TestRSA2WithSHA256_Sign(t *testing.T) {
	rsa2 := &RSAWithSHA256{}
	privateKey := "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC2X977LaTocs9QuSleAx1gibGvMhAFW3bBAf1wr3YP26ckOIqa+FIcET4tiBlCleq+Ci4sRBBdcKPtU0c/f6SlDPP9OqnqRXUEuXsiUMJ7OPrstoIddLCFqHHmaRpXINfZjmgGtKvruDf78bX2FnYa2uRGNhw3ckIOlF4A7NEIiJZl3fum1/qA6oEGTjn1QiB5//8EN3ERYCDBkhgS1g2T0kuL6Ank+0+SWuE0Fm08vU4uP4icb5LRdRHfoerfHcEIxhbVtihWHSmLgGibNEI7dxFiAz3Wotp8zsonwF9Q5cEbpYPpHVpx/i7RpF7yFUm5M61uzonQah7UNrV1Ww4JAgMBAAECggEAf8bJ9tgzCz2tbRReHGU4RvQSTvpXaTl9CZt4U2RL5q5x+5m12wASn2GhW8tYT2O0NXPyh8ckZCNQZy4K5D1tQMrDg+9/Lwl9BFNkJ1XH/QeeHw18OmEQcITlUJbhApybPu1cix44ug23A8mFQKbaFtS4TU0KFfrytz5SYnmJt1y0oOMN+dmYxnbh+w2Av/N09hj61a47WocOQDW8ytBoDD4nZseFUSQ+vSNK3ialvZQ7zXTvyLads6yONcbOCdk8u4+1UUnex3PX+ySNSjQ1JuTQot+FloAZyywOzXFeVkCtZuQpe/gj9ok8i7iuT/UK/5ouDbJJ9PTJxY14t6650QKBgQDZ8TsBzeYcnZsDQC74YVzEsZhT9MdwAOj9LlSy17g3IRT/zHAy4L3alzOjOT8eLYDw3G1oZLo7zniM1L3kePjE2hhnA0yiPBJcOJzslr40nWdTRc8E4ag3Q5Mv0oQXwU9udLnUNak5CXFsV+cx6VcdcGhHzggxY9+rdyrE7K8x3QKBgQDWOKMjbxhPdLt8zd54+Jy2kUstzW0FbO9FTGw3cwHyAj2I6icjIdMug1RKUi+gnQgu8jkVSVRfUqIlJVtCR4dyYomIJIeHk87wGMv84VKugKylCjYxjYgR0DCaXRWiCHQb+7esJ4WXt7Nww7dPfziZylRbB7WPZ+/jhPhHl0mIHQKBgEBqJhCQdJS8mFZLoBZVYH/aJbWawV9/RV2fVfVOAOp6YqSAHiFLf5Gd4us5PkiDFnsaC1QxgUGv8r1dG4rtnklAVLoNpZbFvn93VBoxK6KNaz6XgWpl77v1wwj9ZYFH51w0L8Bi49Mx0U4+ZNzBpLfUw12FrbI7XJ5nKELv2ZAZAoGBAIbRXDJnr3gJ8hjIg3PEmvP3GsY3m54ngaouP4jiE15YdJufKYRdvEdwlXK0qI6/ZTAOd0hjPvtCyRLxoK5kz+R4CTAqNTVpG3pVUMPUlrGF/6FafOLQvMrhKEVtwbiY82HNGDn7IYNrND4Knmokmd2HzXEAuA4Jjpq0y4BawQctAoGBAJZauDM0SONibPROwobeQZPdv5YW4EuQnMBoittmjsqjv2V5csw0K9FJVjmtcOURqKpBw+MJWuxEpV5MNiSqBmBmnyGk6ayRKVFVhkmj8OOzHhdGe+6xJa+L59DyHyqSY45Yu6DBLop37/VzWwqBitcKWRmyQzfF9BPObV1xR7lD"
	data := "app_id=2016080300154586&biz_content={\"out_trade_no\":\"201911270908381158722\",\"total_amount\":\"88.88\",\"subject\":\"iPhone Xs Max 256G\",\"buyer_id\":\"2088102175953034\"}&charset=utf-8&format=json&method=alipay.trade.create&sign_type=RSA2&timestamp=2021-04-22 22:51:34&version=1.0"
	sign := rsa2.Sign(data, privateKey)
	if sign != "EGEEyTFPdagW1obeYAWWO23tz/j7Q0SXnw+mW/p9JGNY4f0ZyTY7CV/nbN3exK9GW3bbINQsp/TwAqfg1tfdjVuyvAQvxw/jnnIzUSTnLPJC9feC6chAcovCVk5SVwW1IYH8pReOcIXzHoTVMXw3kmD+vTBXvLeSy75IQi4vICpl4gJk8YlLFcB68mArkvmBh+74vv8xG4zFFaNOsVNIU5wIJmE54dr5vFJ1GLRHdk5JqIPZO5vptk8OsJPvOBReVs24UeK1TwEjlgEKjEAdF1/waJlvyd4K57502LYqzuu/B2WWxDv3Rh9mrYv2rMB9I14q9i7fcZJtiMHlMAFhtA==" {
		t.Errorf("RSA2????????????????????????")
	}
}

func TestRSA2WithSHA256_Verify(t *testing.T) {
	rsa2 := &RSAWithSHA256{}
	publickKey := `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAh+PU20M0Q4g6GEzM4o3OYfeZFUX3r0p+IxMlTbUal1zIOPCLOm1jemZk+QUaBfRCFsjUA18BTOosBXkXI7LzdUlAGIx9P8q/W2PHqfx53yRdoqKVr6bYu4gH2ryM7baE7D/R7BVp2/CD0j6vUEwSiR8GpSSbEAua5Apc8FvZ8F3OHh/gLUYz0W+VEyrcDrquR2+Xz3R5Rd9Yob1CAwNtzjAgpALkTRTKcQ5EG8p0+UsjxjRoM69xNjtK8gXr61R5LxlyvcCCHSOIHomcyeRnSwUvSmrr3tQ0bSanzBiyOs4sFt+uCcAansMvrn2CxVFYD0JJBQG/Is9AOdeF/f4NCQIDAQAB`
	data := `{"code":"10000","msg":"Success","out_trade_no":"201911270908381158722","qr_code":"https://qr.alipay.com/bax03891zaki07laqrth006b"}`
	sign := `h0+cfA4/8mBX2IBJa8GaZ9yOAvTQskRLRUjqkRCBNlB1/LG3wp/F7YlwrIZ/a147ddfxsQ9PaID3eRWa1vgMS2IXtKJH5slGGsSCEoRX1IyMqhOjzTSUH/9NkbYRNBXAXRdOFezZKggIWL7qLRgaGjJ0aVHR1w7Fl4ubB8Fi1XnJVfH3qRHvgVG2ev1MADJ4dcx/8nCOdhtwUwea3Q5Zt33ytqI66fpv6ij3LDpaN74cVuv+D6ah9j9dZMZvZp4G08E2iTpBZIuBqsIIEVrJxbtQ2Iuqrw9BfJguOMxbcmpAfrjaJjYa3PDqEDapWfjs1Cd1CFOa2+NHkq9BNRNp0w==`
	isOk := rsa2.Verify(data, publickKey, sign)
	if !isOk {
		t.Errorf("RSA2????????????????????????")
	}
}
