import request from '../axios/index'

// 新建产品
export function createProduct(param) {
    return request({
		url: '/product/create',
		method: 'post',
		data: param,
	})
}

// 更新产品
export function updateProduct(param) {
    return request({
		url: '/product/update',
		method: 'put',
		data: param,
	})
}

// 删除产品
export function deleteProduct(param) {
    return request({
		url: '/product/delete',
		method: 'delete',
		data: param,
	})
}

// 查询产品信息
export function queryProductInfo(param) {
    return request({
		url: '/product/info',
		method: 'get',
		params: param,
	})
}

// 查询产品列表
export function queryProductList(param) {
    return request({
		url: '/product/list',
		method: 'get',
		params: param,
	})
}