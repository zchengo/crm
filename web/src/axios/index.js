import axios from 'axios';
import router from '../router/index';
import { message } from 'ant-design-vue';

const host = window.location.hostname

switch (host) {
    case 'www.zocrm.cloud':
        axios.defaults.baseURL = 'https://www.zocrm.cloud/api'
        break;
    default:
        axios.defaults.baseURL = 'http://127.0.0.1:8000/api'
        break;
}

const request = axios.create({
    timeout: 5000,
    headers: {
        'Content-Type': 'application/json;charset=UTF-8'
    }
});

request.interceptors.request.use(config => {
    config.headers['uid'] = localStorage.getItem('uid')
    config.headers['ver'] = localStorage.getItem('ver')
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
    router.push('/error');
    return Promise.reject(error)
})

export default request;