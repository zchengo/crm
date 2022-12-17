import request from '../axios/index'

// 初始化数据
export function initData(param) {
    return request({
		url: '/init/data',
		method: 'post',
		data: param,
	})
}
