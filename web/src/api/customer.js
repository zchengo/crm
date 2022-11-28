import request from '../axios/index'

// 新建客户
export function createCustomer(param) {
    return request({
		url: '/customer/create',
		method: 'post',
		data: param,
	})
}

// 更新客户
export function updateCustomer(param) {
    return request({
		url: '/customer/update',
		method: 'put',
		data: param,
	})
}

// 删除客户
export function deleteCustomer(param) {
    return request({
		url: '/customer/delete',
		method: 'delete',
		data: param,
	})
}

// 查询客户列表
export function queryCustomerList(param) {
    return request({
		url: '/customer/list',
		method: 'get',
		params: param,
	})
}

// 查询客户信息
export function queryCustomerInfo(param) {
    return request({
		url: '/customer/info',
		method: 'get',
		params: param,
	})
}

// 查询客户信息
export function queryCustomerOption(param) {
    return request({
		url: '/customer/option',
		method: 'get',
		params: param,
	})
}