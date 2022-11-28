package common

import (
	"crm/global"
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

// 发送邮件
// QQ邮箱：SMTP 服务器地址：smtp.qq.com（SSL协议端口：465/994 | 非SSL协议端口：25）
// 163邮箱：SMTP 服务器地址：smtp.163.com（端口：25）
func SendMail(email, content string) error {
	smtp := global.Config.Mail.Smtp
	secret := global.Config.Mail.Secret
	sender := global.Config.Mail.Sender
	m := gomail.NewMessage()
	m.SetHeader("From", sender) 		// 发件人
	m.SetHeader("To", email)   			// 收件人，可以多个收件人，但必须使用相同的 SMTP 连接
	m.SetHeader("Cc", email)   			// 抄送，可以多个
	m.SetHeader("Bcc", email)  			// 暗送，可以多个
	m.SetHeader("Subject", "ZOCRM")     // 邮件主题
	m.SetBody("text/html", content)
	d := gomail.NewDialer(smtp, 25, sender, secret)
	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("qq mail send error : %s", err)
		return err
	}
	return nil
}