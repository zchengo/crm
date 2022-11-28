import request from '../axios/index'

// 获取数据汇总
export function getSummary(param) {
    return request({
		url: '/dashboard/sum',
		method: 'get',
		params: param,
	})
}
