# crm

<a href="#公众号"><img src="https://img.shields.io/badge/公众号-GoCode-%2302DF6D" /></a>
<a href="https://zocrm.cloud"><img src="https://img.shields.io/badge/官网-zocrm.cloud-%230092FF" /></a>

## 项目简介

客户关系管理系统，基于 Vue + Go 实现，主要功能有仪表盘、客户管理、合同管理、产品管理，订阅等功能。

在线演示：https://zocrm.cloud

开发文档：https://docs.zocrm.cloud

## 技术选型

### 前端技术

| 技术 | 说明 | 相关文档 |
|---|---|---|
| Vue.js | 前端框架 | https://v3.cn.vuejs.org |
| Vue Router | 页面路由 | https://router.vuejs.org |
| Axios | 网络请求库 | https://axios-http.com |
| Pinia | 状态管理 | https://pinia.vuejs.org |
| Vite | 构建工具 | https://vitejs.cn |
| Ant Design Vue | 前端UI组件库 | https://www.antdv.com |
| Apache ECharts | 可视化图表库 | https://echarts.apache.org |
| Moment | 日期库 | https://momentjs.com |

### 后端技术

| 技术 | 说明 | 相关文档 |
|---|---|---|
| Gin | Web框架 | https://gin-gonic.com |
| Gorm | ORM框架 | https://gorm.io |
| Jwt | 用户认证 | https://github.com/golang-jwt/jwt |
| Viper | 配置管理 | https://github.com/spf13/viper |
| Redis | 数据缓存 | https://github.com/go-redis/redis |
| Mail | 邮件服务SDK | https://github.com/go-gomail/gomail |

## 快速运行

系统运行环境：

| 环境 | 版本 | 下载地址 |
|---|---|---|
| go | >= 1.19.2 | https://golang.google.cn/dl |
| mysql | >= 8.0.31 | https://www.mysql.com/downloads |
| redis | >= 7.0.5 | https://redis.io/download |
| node | >= 18.12.0 | https://nodejs.org/en/download |


在终端中，执行如下命令，初始化并运行。

```
$ cd server
$ go mod tidy
$ go build -o server main.go (windows编译命令为 go build -o server.exe main.go )

# 运行二进制
$ ./server (windows运行命令为 server.exe)

$ cd web
$ npm install
$ npm run dev
```

项目运行成功后，打开浏览器访问 http://127.0.0.1:8060

## 公众号

关注公众号「**GoCode**」，回复「**部署指南**」，获取详细部署指南！

![公众号图片](https://zocrm.cloud/gzh_qrcode.jpg)

## 许可证

[MIT License](https://github.com/zchengo/crm/blob/main/LICENSE) 

Copyright (c) 2022 zchengo 