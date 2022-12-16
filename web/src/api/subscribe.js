import request from '../axios/index'

// 订阅专业版
export function subscribePay(param) {
    return request({
		url: '/subscribe/pay',
		method: 'post',
		data: param,
	})
}

// 获取订阅信息
export function getSubscribeInfo(param) {
    return request({
		url: '/subscribe/info',
		method: 'get',
		params: param,
	})
}

