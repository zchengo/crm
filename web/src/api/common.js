import request from '../axios/index'

// 初始化数据库
export function initDatabase(param) {
    return request({
		url: '/common/database/init',
		method: 'post',
		data: param,
	})
}

// 文件移除
export function fileRemove(param) {
    return request({
		url: '/common/file/remove',
		method: 'delete',
		data: param,
	})
}
