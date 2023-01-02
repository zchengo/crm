package service

import (
	"crm/global"
	"crm/models"
	"crm/response"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/smartwalle/alipay/v3"
	"github.com/smartwalle/xid"
)

const (
	Paying = 1 // 支付中
	Payed  = 2 // 已支付
)

type SubscribeService struct {
	noticeService *NoticeService
}

func NewSubscribeService() *SubscribeService {
	subscribeService := SubscribeService{
		noticeService: &NoticeService{},
	}
	return &subscribeService
}

// 订阅专业版，发起支付
func (s *SubscribeService) Pay(param models.SubscribePayParam) (*models.SubscribePayUrl, int) {

	// 构建订单支付信息
	tradeNo := fmt.Sprintf("%d", xid.Next())
	var p = alipay.TradePagePay{}
	p.NotifyURL = global.Config.Alipay.NotifyURL
	p.ReturnURL = global.Config.Alipay.ReturnURL
	fmt.Println(p.ReturnURL)
	p.Subject = "支付测试:" + tradeNo
	p.OutTradeNo = tradeNo
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	switch param.Version {
	case 2:
		p.TotalAmount = "18.00"
	case 3:
		p.TotalAmount = "198.00"
	}

	// 缓存订单信息
	order, _ := json.Marshal(&models.SubscribePayOrder{
		Uid:  param.Uid,
		Version: param.Version,
	})
	err := global.Rdb.Set(ctx, tradeNo, string(order), time.Minute*30).Err()
	if err != nil {
		return nil, response.ErrCodeFailed
	}

	// 设置支付状态
	key := fmt.Sprintf("uid:%v:pay:status", param.Uid)
	if err := global.Rdb.Set(ctx, key, Paying, time.Minute*10).Err(); err != nil {
		return nil, response.ErrCodeFailed
	}

	// 返回支付链接
	payUrl, err := global.Alipay.TradePagePay(p)
	if err != nil {
		return nil, response.ErrCodeFailed
	}
	subscribePayUrl := models.SubscribePayUrl{
		PayUrl: payUrl.String(),
	}
	return &subscribePayUrl, response.ErrCodeSuccess
}

// 支付成功回调
func (s *SubscribeService) Callback(outTradeNo string) (string, int) {

	// 验证订单信息
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo
	rsp, err := global.Alipay.TradeQuery(p)
	if err != nil {
		log.Printf("验证订单 %s 信息发生错误: %s", outTradeNo, err.Error())
		return StringNull, response.ErrCodeFailed
	}
	if !rsp.IsSuccess() {
		log.Printf("验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Content.Msg, rsp.Content.SubMsg)
		return StringNull, response.ErrCodeFailed
	}

	// 获取订单信息
	var order models.SubscribePayOrder
	orderJson, err := global.Rdb.Get(ctx, outTradeNo).Result()
	if err != nil {
		return StringNull, response.ErrCodeFailed
	}
	if err := json.Unmarshal([]byte(orderJson), &order); err != nil {
		return StringNull, response.ErrCodeFailed
	}

	// 创建订阅信息
	var expired int64
	var content string
	switch order.Version {
	case 2:
		expired = time.Now().Unix() + int64(2592000)
		content = SUBSCRIBE_NOTICE_TEMPLATE1
	case 3:
		expired = time.Now().Unix() + int64(31104000)
		content = SUBSCRIBE_NOTICE_TEMPLATE2
	}
	subscribe := models.Subscribe{
		Version: order.Version,
		Expired: expired,
	}
	if global.Db.Table(SUBSCRIBE).Where("uid = ?", order.Uid).First(&models.Subscribe{}).RowsAffected == 0 {
		subscribe.Uid = order.Uid
		subscribe.Created = time.Now().Unix()
		err := global.Db.Table(SUBSCRIBE).Create(&subscribe).Error
		if err != nil {
			return StringNull, response.ErrCodeFailed
		}
	} else {
		subscribe.Updated = time.Now().Unix()
		err = global.Db.Model(&models.Subscribe{}).Where("uid = ?", order.Uid).Updates(&subscribe).Error
		if err != nil {
			return StringNull, response.ErrCodeFailed
		}
	}

	// 订阅通知
	s.noticeService.Create(&models.NoticeParam{
		Content: content,
		Creator: order.Uid,
	})

	return global.Config.Alipay.PaySuccessURL, response.ErrCodeSuccess
}

// 异步验签
func (s *SubscribeService) Notify(data url.Values, outTradeNo string) int {
	ok, err := global.Alipay.VerifySign(data)
	if err != nil {
		log.Println("异步通知验证签名发生错误", err)
		return response.ErrCodeFailed
	}

	if !ok {
		log.Println("异步通知验证签名未通过")
		return response.ErrCodeFailed
	}

	log.Println("异步通知验证签名通过")

	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo
	rsp, err := global.Alipay.TradeQuery(p)
	if err != nil {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s \n", outTradeNo, err.Error())
		return response.ErrCodeFailed
	}
	if !rsp.IsSuccess() {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s-%s \n", outTradeNo, rsp.Content.Msg, rsp.Content.SubMsg)
		return response.ErrCodeFailed
	}

	log.Printf("订单 %s 支付成功 \n", outTradeNo)
	return response.ErrCodeSuccess
}

// 获取订阅信息
func (s *SubscribeService) GetInfo(uid int64) (*models.SubscribeInfo, int) {
	var si models.SubscribeInfo
	if err := global.Db.Table(SUBSCRIBE).First(&si, "uid = ?", uid).Error; err != nil {
		return nil, response.ErrCodeFailed
	}
	// 判断用户订阅是否过期
	if si.Version == 2 && time.Now().Unix() > int64(si.Expired) {
		err := global.Db.Model(&models.Subscribe{}).Where("uid = ?", uid).Update("version", 1).Error
		if err != nil {
			return nil, response.ErrCodeFailed
		}
	}
	if err := global.Db.Table(SUBSCRIBE).First(&si, "uid = ?", uid).Error; err != nil {
		return nil, response.ErrCodeFailed
	}
	return &si, response.ErrCodeSuccess
}
