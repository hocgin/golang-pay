package allin

import (
	"fmt"
	"github.com/hocgin/golang-pay/alipay"
	"testing"
)

func TestConfigs(t *testing.T) {
	service := new(AllInPayServiceImpl)
	alipayService := alipay.CreateConfigStorage("", "", "").CreateService()

	service.Configs(alipayService)
	payResult, _ := service.DoPay(TradeInfo{})
	fmt.Sprintln(payResult)
}
