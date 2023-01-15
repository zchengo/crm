# 页面中的加载进度条

Crm 系统使用 NProgress 作为页面中的加载进度条。

## 什么是 NProgress？

NProgress 是一个适用于应用程序的超薄进度条，灵感来自谷歌、YouTube和媒体。了解有关 NProgress 的更多信息，请访问[NProgress](https://github.com/rstacruz/nprogress)。

## 安装 NProgress

添加 ```nprogress.js``` 和 ```nprogress.css``` 到您的项目。

```html
<script src='nprogress.js'></script>
<link rel='stylesheet' href='nprogress.css'/>
```

或者使用 ```npm``` 安装：

```bash
$ npm install --save nprogress
```

## 在 Vue Router 中使用

在 Vue Router 中使用 nprogress，请参考如下：

```js
const router = createRouter({
  history: createWebHashHistory(), routes
})

NProgress.configure({ easing: 'ease', speed: 500, showSpinner: false });

router.beforeEach((to, from, next) => {
  NProgress.start() // 进度条开始
  next()
})

router.afterEach(() => {
  NProgress.done() // 进度条结束
})
```

查看完整代码，请访问[crm/web/src/router/index.js](https://github.com/zchengo/crm/blob/main/web/src/router/index.js)。

想要了解有关 NProgress 的更多信息，请访问[NProgress](https://github.com/rstacruz/nprogress)。







