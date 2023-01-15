# 封装 axios 请求库

## 什么是 axios？

axios 是浏览器和 node.js 的一个简单的基于 promise 的 HTTP 客户端。axios 在一个具有非常可扩展界面的小软件包中提供了一个简单易用的库。

了解有关 axios 的更多信息，请访问[axios](https://github.com/axios/axios)。

## 封装 axios

封装 axios 请求库，请参考如下：

```js
import axios from 'axios';
import router from '../router/index';
import { message } from 'ant-design-vue';

const host = window.location.hostname

switch (host) {
    case 'zocrm.cloud':
        axios.defaults.baseURL = 'https://zocrm.cloud/api'
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
```

### 初始化请求的 baseURL

```js
const host = window.location.hostname

switch (host) {
    case 'zocrm.cloud':
        axios.defaults.baseURL = 'https://zocrm.cloud/api'
        break;
    default:
        axios.defaults.baseURL = 'http://127.0.0.1:8000/api'
        break;
}
```

通过 ```window.location.hostname``` 获取主机名（不包含端口号），主机名指的是 ```IP地址``` 或 ```域名```，然后根据不同的主机名，设置不同的 ```axios.defaults.baseURL```。

### 设置请求超时时间与请求头

```js
const request = axios.create({
    timeout: 5000,
    headers: {
        'Content-Type': 'application/json;charset=UTF-8'
    }
});
```

### 请求拦截与响应拦截

请求拦截，请参考如下：

```js
request.interceptors.request.use(config => {
    config.headers['uid'] = localStorage.getItem('uid')
    config.headers['token'] = localStorage.getItem('token')
    return config
})
```

前端每次发送请求时，都会被 axios 拦截。axios 会先给这个请求设置请求头 ```uid``` 和 ```token```，然后再向服务端发送请求。

响应拦截，请参考如下：

```js
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
```

后端响应的数据会先被 axios 拦截，如果响应的结果是正确的，会直接返回响应结果，然后将数据渲染到指定的页面；如果响应的结果是错误的，会路由到错误页面。

了解有关 axios 请求拦截与响应拦截的更多信息，请访问[axios interceptors](https://github.com/axios/axios#interceptors)。




