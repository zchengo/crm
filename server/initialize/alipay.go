package initialize

import (
	"crm/global"
	"log"

	"github.com/smartwalle/alipay/v3"
)

func Alipay() {
	var err error
	appId := global.Config.Alipay.AppId
	privateKey := global.Config.Alipay.PrivateKey
	global.Alipay, err = alipay.New(appId, privateKey, false);
	if err != nil {
		log.Println("初始化支付宝支付服务失败", err)
		return
	}

	// 加载支付宝证书
	if err = global.Alipay.LoadAppPublicCertFromFile(global.Config.Alipay.AppPublicCert); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}

	if err = global.Alipay.LoadAliPayRootCertFromFile(global.Config.Alipay.AlipayRootCert); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
	if err = global.Alipay.LoadAliPayPublicCertFromFile(global.Config.Alipay.AlipayPublicCert); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
}