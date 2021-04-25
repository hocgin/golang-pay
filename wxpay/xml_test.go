package wxpay

import (
	"encoding/xml"
	"fmt"
	"github.com/hocgin/golang-pay/core/ops"
	"testing"
)

func TestRefundQuery(t *testing.T) {
	xmlStr := `<xml>
<total_refund_count>12</total_refund_count>
<refund_count>2</refund_count>
<out_refund_no_0>coupon_refund_count_0</out_refund_no_0>
<coupon_refund_count_0>1</coupon_refund_count_0>
<coupon_refund_id_0_0>coupon_refund_count_1</coupon_refund_id_0_0>
</xml>`
	var result interface{}
	result = &RefundQueryResponse{}
	_ = xml.Unmarshal([]byte(xmlStr), result)
	if response, isOk := result.(ops.SetResponseBody); isOk {
		response.SetResponseBody(xmlStr)
	}
	if response, isOk := result.(ops.AfterPropertiesSet); isOk {
		response.AfterPropertiesSet()
	}
	fmt.Print(result)
}
