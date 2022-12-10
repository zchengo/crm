import request from '../axios/index'

// 新建合同
export function createContract(param) {
    return request({
		url: '/contract/create',
		method: 'post',
		data: param,
	})
}

// 更新合同
export function updateContract(param) {
    return request({
		url: '/contract/update',
		method: 'put',
		data: param,
	})
}

// 删除合同
export function deleteContract(param) {
    return request({
		url: '/contract/delete',
		method: 'delete',
		data: param,
	})
}

// 查询合同列表
export function queryContractList(param) {
    return request({
		url: '/contract/list',
		method: 'get',
		params: param,
	})
}

// 查询合同信息
export function queryContractInfo(param) {
    return request({
		url: '/contract/info',
		method: 'get',
		params: param,
	})
}

// 查询添加的产品列表
export function queryContractPlist(param) {
    return request({
		url: '/contract/plist',
		method: 'post',
		data: param,
	})
}