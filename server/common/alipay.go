package common

import (
	"context"
	"crm/global"
	"net/http"
	"strconv"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
)

type Alipay struct {
}

func GetAlipay() *Alipay {
	return &Alipay{}
}

func (a *Alipay) PagePay(tradeNo string, totalAmount float64) (string, error) {
	bm := make(gopay.BodyMap)
	bm.Set("subject", "服务订阅").
		Set("out_trade_no", tradeNo).
		Set("total_amount", totalAmount).
		Set("timeout_express", "2m")

	payUrl, err := global.Alipay.TradePagePay(context.Background(), bm)

	if err != nil {
		if bizErr, ok := alipay.IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			return "", err
		}
		xlog.Errorf("client.TradePay(%+v),err:%+v", bm, err)
		return "", err
	}
	return payUrl, nil
}

func (a *Alipay) VerifySign(req *http.Request) gopay.BodyMap {
	notifyReq, err := alipay.ParseNotifyToBodyMap(req)
	if err != nil {
		xlog.Error(err)
		return nil
	}

	// 支付宝异步通知验签（公钥模式）
	if _, err = alipay.VerifySignWithCert(global.Config.Alipay.AlipayPublicCert, notifyReq); err != nil {
		xlog.Error(err)
		return nil
	}
	return notifyReq
}

func (a *Alipay) GenTradeNo() string {
	return time.Now().Format("20060102150405") + strconv.Itoa(RandInt(100000, 999999))
}
