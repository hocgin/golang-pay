package alipay

import (
	"errors"
	"github.com/hocgin/golang-pay/allin"
	"github.com/hocgin/golang-pay/core"
	"time"
)

func (this AliPayPaymentService) DoPay(trade allin.TradeInfo) (*allin.PayResult, error) {
	result := new(allin.PayResult)
	result.PayWay = trade.PayWay
	switch trade.PayWay {
	case allin.AliPayApp:
		req := TradeAppPayRequest{
			BizContent: TradeAppPayBizContent{OutTradeNo: trade.OutTradeNo, Subject: trade.Subject, TotalAmount: trade.TotalAmount},
		}
		req.NotifyUrl = trade.NotifyUrl
		if rep, err := this.TradeAppPay(req); err != nil {
			return nil, err
		} else {
			result.Response = rep
			result.Result = rep
		}
	case allin.AliPayWeb:
		req := TradePagePayRequest{
			BizContent: TradePagePayBizContent{OutTradeNo: trade.OutTradeNo, Subject: trade.Subject, TotalAmount: trade.TotalAmount,
				ProductCode: "FAST_INSTANT_TRADE_PAY"},
		}
		req.NotifyUrl = trade.NotifyUrl
		if rep, err := this.TradePagePay(req); err != nil {
			return nil, err
		} else {
			result.Response = rep
			result.Result = rep
		}
	case allin.AliPayNative:
		req := TradeWapPayRequest{
			BizContent: TradeWapPayBizContent{OutTradeNo: trade.OutTradeNo, Subject: trade.Subject, TotalAmount: trade.TotalAmount,
				ProductCode: "QUICK_WAP_WAY",
				QuitUrl:     trade.QuitUrl},
		}
		req.NotifyUrl = trade.NotifyUrl
		if rep, err := this.TradeWapPay(req); err != nil {
			return nil, err
		} else {
			result.Response = rep
			result.Result = rep
		}
	case allin.AliPayQrCode:
		req := TradePreCreateRequest{
			BizContent: TradePreCreateBizContent{OutTradeNo: trade.OutTradeNo, Subject: trade.Subject, TotalAmount: trade.TotalAmount},
		}
		req.NotifyUrl = trade.NotifyUrl
		if rep, err := this.TradePreCreate(req); err != nil {
			return nil, err
		} else {
			result.Response = rep
			result.Result = rep.QrCode
		}
	default:
		return nil, errors.New("alipay not support the pay way")
	}
	return result, nil
}
func (this AliPayPaymentService) DoRefund(refund allin.RefundInfo) error {
	_, err := this.TradeRefund(TradeRefundRequest{
		BizContent: TradeRefundBizContent{OutTradeNo: refund.OutTradeNo, RefundAmount: refund.RefundAmount},
	})
	return err
}
func (this AliPayPaymentService) DoQueryPay(outTradeNo string) (*allin.QueryPayResult, error) {
	_, err := this.TradeQuery(TradeQueryRequest{
		BizContent: TradeQueryBizContent{OutTradeNo: outTradeNo},
	})
	return nil, err
}
func (this AliPayPaymentService) DoQueryRefund(outTradeNo string) (*allin.QueryRefundResult, error) {
	_, err := this.TradeQuery(TradeFastpayRefundQueryRequest{
		BizContent: TradeFastpayRefundQueryBizContent{OutTradeNo: outTradeNo},
	})
	return nil, err
}
func (this AliPayPaymentService) DoPayMessage(queryStrings string) (*allin.PayMessageResult, error) {
	result := new(allin.PayMessageResult)
	result.Channel = core.AliPay
	rep, err := this.TradeStatusSyncMessage(queryStrings)
	if rep != nil {
		result.PayAmount = rep.BuyerPayAmount
		result.TotalAmount = rep.TotalAmount
		result.OutTradeNo = rep.OutTradeNo
		result.TradeNo = rep.TradeNo
		result.TradeStatus = rep.TradeStatus
		result.BuyerId = rep.BuyerId
		// rep.NotifyTime
		result.NotifyTime = time.Now()
		result.Response = rep
	}

	return result, err
}
func (this AliPayPaymentService) DoRefundMessage(queryStrings string) (*allin.RefundMessageResult, error) {
	_, err := this.TradeStatusSyncMessage(queryStrings)
	return nil, err
}
