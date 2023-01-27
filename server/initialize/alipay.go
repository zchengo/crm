package initialize

import (
	"crm/global"
	"log"

	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
)

func Alipay() {
	pay := global.Config.Alipay
	client, err := alipay.NewClient(pay.AppId, pay.PrivateKey, false)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 设置支付宝请求、公钥证书模式
	client.SetReturnUrl(pay.ReturnURL).SetNotifyUrl(pay.NotifyURL)
	err = client.SetCertSnByPath(pay.AppPublicCert, pay.AlipayRootCert, pay.AlipayPublicCert)
	if err != nil {
		log.Printf("init alipay cert error: %s", err)
		return
	}
	global.Alipay = client
}
