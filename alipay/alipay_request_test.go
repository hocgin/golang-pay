package alipay

import (
	"fmt"
	"testing"
)

var (
	alipayConfig = CreateConfigStorage("2016080300154586",
		"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",
		IsDev(true))
	service = alipayConfig.CreateService()
)

func TestTradePay(t *testing.T) {
	request := &TradePayRequest{
		BizContent: TradePayBizContent{
			OutTradeNo: "201911270908381158722",
		},
	}
	response, _ := service.TradePay(request)
	fmt.Print(response)
}

func TestTradeCreate(t *testing.T) {
	request := &TradeCreateRequest{
		BizContent: TradeCreateBizContent{
			Subject:     "????????????",
			OutTradeNo:  "201911270908381158722",
			TotalAmount: 1.0,
			BuyerId:     "2088102170344093",
		},
	}
	response, _ := service.TradeCreate(request)
	fmt.Print(response)
}

func TestTradePreCreate(t *testing.T) {
	request := &TradePreCreateRequest{
		BizContent: TradePreCreateBizContent{
			Subject:     "????????????",
			OutTradeNo:  "201911270908381158722",
			TotalAmount: 1.0,
		},
	}
	response, _ := service.TradePreCreate(request)
	fmt.Print(response)
}

func TestTradeAppPay(t *testing.T) {
	request := &TradeAppPayRequest{
		BizContent: TradeAppPayBizContent{
			Subject:     "????????????",
			OutTradeNo:  "201911270908381158722",
			TotalAmount: 1.0,
		},
	}
	response, _ := service.TradeAppPay(request)
	fmt.Println(response)
}

func TestTradePagePay(t *testing.T) {
	request := &TradePagePayRequest{
		BizContent: TradePagePayBizContent{
			OutTradeNo:  "201911270908381158722",
			TotalAmount: 1.0,
		},
	}
	response, _ := service.TradePagePay(request)
	fmt.Println(response)
}
