package common

import (
	"crm/global"
	"crm/models"
	"crypto/tls"
	"fmt"
	"log"

	"gopkg.in/gomail.v2"
)

// 发送邮件（系统级别）
// QQ邮箱：SMTP 服务器地址：smtp.qq.com（SSL协议端口：465/994 | 非SSL协议端口：25）
// 163邮箱：SMTP 服务器地址：smtp.163.com（端口：25）
func SendMail(email, content string) error {
	smtp := global.Config.Mail.Smtp
	secret := global.Config.Mail.Secret
	sender := global.Config.Mail.Sender
	m := gomail.NewMessage()
	m.SetHeader("From", sender)     // 发件人
	m.SetHeader("To", email)        // 收件人，可以多个收件人，但必须使用相同的 SMTP 连接
	m.SetHeader("Cc", email)        // 抄送，可以多个
	m.SetHeader("Bcc", email)       // 暗送，可以多个
	m.SetHeader("Subject", "ZOCRM") // 邮件主题
	m.SetBody("text/html", content)
	d := gomail.NewDialer(smtp, 465, sender, secret)
	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("qq mail send error : %s", err)
		return err
	}
	return nil
}

// 发送邮件给客户（客户级别）
func SendMailToCustomer(mp models.MailParam) error {
	m := gomail.NewMessage()
	m.SetHeader("From", mp.Sender)
	m.SetHeader("To", mp.Receiver)
	m.SetHeader("Subject", mp.Subject)
	m.SetBody("text/html", mp.Content)
	if mp.Attachment != "" {
		m.Attach(mp.Attachment)
	}
	d := gomail.NewDialer(mp.Smtp, mp.Port, mp.Sender, mp.AuthCode)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		log.Printf("send mail to customer error : %s", err)
		return err
	}
	return nil
}

// 检测STMP服务是否可连接
func DialMail(smtp string, port int, sender, authCode string) error {
	d := gomail.NewDialer(smtp, port, sender, authCode)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	_, err := d.Dial()
	return err
}
