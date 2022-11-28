# crm

客户关系管理系统，基于 Vue + Go 实现，主要功能有仪表盘、客户管理、合同管理、产品管理，订阅等功能。

# 项目演示

开源ZOCRM官网：https://www.zocrm.cloud

# 项目文档

- [技术选型](https://github.com/zchengo/crm/blob/main/docs/techstack.md)
- [目录结构](https://github.com/zchengo/crm/blob/main/docs/structure.md)

# 快速运行

系统运行环境：

| 环境 | 版本 | 下载地址 |
|---|---|---|
| go | >= 1.19.2 | https://golang.google.cn/dl/ |
| mysql | >= 8.0.31 | https://www.mysql.com/downloads/ |
| redis | >= 7.0.5 | https://redis.io/download/ |
| node | >= 18.12.0 | https://nodejs.org/en/download/ |

在终端（Terminal）中，执行如下命令，进行项目的初始化和运行。

```
$ git clone https://github.com/zchengo/crm.git

$ cd server
$ go mod tidy
$ go build -o server main.go (windows编译命令为 go build -o server.exe main.go )

# 运行二进制
$ ./server (windows运行命令为 server.exe)

$ cd web
$ npm install
$ npm run dev
```

运行成功后，使用浏览器访问：http://127.0.0.1:8008/#/login

账号：12345@qq.com 密码：123

# 问题反馈

你可以提交 Issues ，也可以通过 [知乎](https://www.zhihu.com/people/87-4-8-5) 或 [CSDN](https://blog.csdn.net/m0_47890251?spm=1000.2115.3001.5343) 联系作者。

# 许可证

[MIT License](https://github.com/zchengo/crm/blob/main/LICENSE) 