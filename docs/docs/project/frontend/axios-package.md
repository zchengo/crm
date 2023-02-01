# 网络请求库封装

Crm 系统采用 axios 作为网络请求库。

## 什么是 axios ？

axios 是浏览器和 node.js 的一个简单的基于 promise 的 HTTP 客户端。axios 在一个具有非常可扩展界面的小软件包中提供了一个简单易用的库。

了解有关 axios 的更多信息，请访问[axios](https://github.com/axios/axios)。

## 封装 axios

封装 axios 请求库，请参考如下：

```js
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
```

### 初始化请求的 baseURL

```js
axios.defaults.baseURL = import.meta.env.VITE_API_BASE_URL
```

通过 ```import.meta.env.VITE_API_BASE_URL``` 获取环境变量。

### 设置请求超时时间与请求头

```js
const request = axios.create({
    // timeout: 5000,`
    headers: {
        'Content-Type': 'application/json;charset=UTF-8'
    }
})
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
    return Promise.reject(error)
})
```

了解有关 axios 请求拦截与响应拦截的更多信息，请访问[axios interceptors](https://github.com/axios/axios#interceptors)。




