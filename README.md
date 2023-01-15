<p align="center">
  <a href="https://zocrm.cloud">
    <img src="https://docs.zocrm.cloud/logo.svg" width="150" height="150" style="filter: drop-shadow(2px 2px 6px #79bbff);"/>
  </a>
</p>

<h1 align="center">
  <a href="https://docs.zocrm.cloud" target="_blank">CRM</a>
</h1>

<div align="center">

一个免费、开源的客户关系管理系统。

<a href="#公众号"><img src="https://img.shields.io/badge/公众号-GoCode-%2302DF6D" /></a>
<a href="https://zocrm.cloud"><img src="https://img.shields.io/badge/在线演示-zocrm.cloud-%230092FF" /></a>
<a href="https://docs.zocrm.cloud"><img src="https://img.shields.io/badge/项目文档-docs.zocrm.cloud-%230092FF" /></a>

</div>

## 简介

客户关系管理系统，基于 Vue + Go 实现，主要功能有仪表盘、客户管理、合同管理、产品管理，订阅等功能。

- 在线演示：[zocrm.cloud](https://zocrm.cloud)

- 项目文档：[docs.zocrm.cloud](https://docs.zocrm.cloud)

## 快速开始

系统运行环境：

| 环境 | 版本 | 下载地址 |
|---|---|---|
| go | >= 1.19.2 | https://golang.google.cn/dl |
| mysql | >= 8.0.31 | https://www.mysql.com/downloads |
| redis | >= 7.0.5 | https://redis.io/download |
| node | >= 18.12.0 | https://nodejs.org/en/download |

在终端中，执行如下命令：

```bash
$ cd server
$ go mod tidy
$ go build -o server main.go (windows编译命令为 go build -o server.exe main.go )

# 运行二进制
$ ./server (windows运行命令为 server.exe)

$ cd web
$ npm install
$ npm run dev
```

初始化和启动成功后，打开浏览器访问[http://127.0.0.1:8060](http://127.0.0.1:8060)。

## 项目文档

想要了解有关项目的更多信息，请访问[docs.zocrm.cloud](https://docs.zocrm.cloud)。

## 公众号

欢迎关注公众号「GoCode」，本公众号专注Go语言技术分享！

<img src="https://docs.zocrm.cloud/gzh_qrcode.jpg" style="border-radius: 8px;" alt="gzh_qrcode"/>

## 许可证

[MIT License](https://github.com/zchengo/crm/blob/main/LICENSE) 

Copyright (c) 2022 zchengo 