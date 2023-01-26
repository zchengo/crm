import request from '../axios/index'

// 保存邮件服务配置
export function saveMailConfig(param) {
    return request({
		url: '/config/save',
		method: 'put',
		data: param,
	})
}

// 删除邮件服务配置
export function deleteMailConfig(param) {
    return request({
		url: '/config/delete',
		method: 'delete',
		data: param,
	})
}

// 开启或关闭邮件服务
export function updateMailConfigStmpStatus(param) {
    return request({
		url: '/config/status',
		method: 'put',
		data: param,
	})
}

// 获取邮件服务配置信息
export function getMailConfig(param) {
    return request({
		url: '/config/info',
		method: 'get',
		data: param,
	})
}

// 检查邮件配置的有效性
export function checkMailConfig(param) {
    return request({
		url: '/config/check',
		method: 'get',
		data: param,
	})
}