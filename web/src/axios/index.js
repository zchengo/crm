import axios from 'axios';
import { message } from 'ant-design-vue';

axios.defaults.baseURL = import.meta.env.VITE_API_BASE_URL

const request = axios.create({
    // timeout: 5000,`
    headers: {
        'Content-Type': 'application/json;charset=UTF-8'
    }
})

request.interceptors.request.use(config => {
    config.headers['uid'] = localStorage.getItem('uid')
    config.headers['token'] = localStorage.getItem('token')
    return config
})

request.interceptors.response.use(response => {
    console.log(response)
    if (response.data.code == 1) {
        message.error('服务器异常！')
    }
    return response;
}, error => {
    console.log(error)
    return Promise.reject(error)
})

export default request;