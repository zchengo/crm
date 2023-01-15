# 快速开始

## 安装

如果你的系统中已经安装了 ```Git```，你可以使用如下方式安装：

```bash
git clone https://github.com/zchengo/crm.git
```

你也可以通过下载压缩包的方式安装[Download ZIP](https://github.com/zchengo/crm/archive/refs/heads/main.zip)。

## 部署

你需要在你的系统中，安装如下运行环境：

| 环境 | 版本 | 下载地址 |
|---|---|---|
| go | >= 1.19.2 | https://golang.google.cn/dl |
| mysql | >= 8.0.31 | https://www.mysql.com/downloads |
| redis | >= 7.0.5 | https://redis.io/download |
| node | >= 18.12.0 | https://nodejs.org/en/download |

::: tip 初始化和运行的前提条件
Go 环境正常、Node 环境正常、MySQL 和 Redis 能正常启动。
:::

然后在终端中，执行如下命令：

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

项目初始化并运行成功后，打开浏览器访问http://127.0.0.1:8060。

详细部署，请参考[部署指南](/project/docs/deploy-guide)。