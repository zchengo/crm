package service

import (
	"crm/common"
	"crm/dao"
	"crm/models"
	"crm/response"
	"encoding/json"
	"time"
)

type SubscribeService struct {
	subscribeDao *dao.SubscribeDao
	noticeDao    *dao.NoticeDao
}

func NewSubscribeService() *SubscribeService {
	subscribeService := SubscribeService{
		subscribeDao: dao.NewSubscribeDao(),
		noticeDao:    dao.NewNoticeDao(),
	}
	return &subscribeService
}

// 订阅专业版，发起支付
func (s *SubscribeService) Pay(param models.SubscribePayParam) (*models.SubscribePayUrl, int) {

	// 构建订单支付信息
	tradeNo := common.GetAlipay().GenTradeNo()
	totalAmount := float64(param.Duration) * 0.6
	payUrl, err := common.GetAlipay().PagePay(tradeNo, totalAmount)
	if err != nil {
		return nil, response.ErrCodeFailed
	}

	order := models.SubscribePayOrder{
		Uid:      param.Uid,
		TradeNo:  tradeNo,
		Duration: param.Duration,
	}

	if err := s.subscribeDao.SetOrder(tradeNo, &order); err != nil {
		return nil, response.ErrCodeFailed
	}

	subscribePayUrl := models.SubscribePayUrl{
		PayUrl: payUrl,
	}
	return &subscribePayUrl, response.ErrCodeSuccess
}

// 支付成功回调
func (s *SubscribeService) PayBack(outTradeNo string) int {

	// 获取订单信息
	var order models.SubscribePayOrder
	orderJson, err := s.subscribeDao.GetOrder(outTradeNo)
	if err != nil {
		return response.ErrCodeFailed
	}
	if err := json.Unmarshal([]byte(orderJson), &order); err != nil {
		return response.ErrCodeFailed
	}

	// 创建订阅信息
	duration := order.Duration * 24 * 60 * 60
	if !s.subscribeDao.IsExists(order.Uid) {
		subscribe := models.SubscribeCreateParam{
			Uid:     order.Uid,
			Version: 2,
			Expired: time.Now().Unix() + duration,
		}
		if err := s.subscribeDao.Create(&subscribe); err != nil {
			return response.ErrCodeFailed
		}
	} else {
		si, err := s.subscribeDao.GetInfo(order.Uid)
		if err != nil {
			return response.ErrCodeFailed
		}
		subscribe := models.SubscribeUpdateParam{
			Uid:     order.Uid,
			Version: 2,
			Expired: si.Expired + duration,
		}
		if err := s.subscribeDao.Update(&subscribe); err != nil {
			return response.ErrCodeFailed
		}
	}

	// 订阅通知
	s.noticeDao.Create(&models.NoticeCreateParam{
		Content: SUBSCRIBE_NOTICE_TEMPLATE1,
		Creator: order.Uid,
	})

	return response.ErrCodeSuccess
}

// 获取订阅信息
func (s *SubscribeService) GetInfo(uid int64) (*models.SubscribeInfo, int) {
	si, err := s.subscribeDao.GetInfo(uid)
	if err != nil {
		return nil, response.ErrCodeFailed
	}
	// 判断用户订阅是否过期
	if si.Version == 2 && time.Now().Unix() > int64(si.Expired) {
		if err := s.subscribeDao.UpdateVersion(uid, 1); err != nil {
			return nil, response.ErrCodeFailed
		}
	}
	subscribeInfo, err := s.subscribeDao.GetInfo(uid)
	if err != nil {
		return nil, response.ErrCodeFailed
	}
	return subscribeInfo, response.ErrCodeSuccess
}
