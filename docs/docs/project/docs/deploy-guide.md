# 部署指南

::: tip 提示
请先安装运行环境，再运行服务端和Web端。
:::

## 环境安装

在您自己的电脑上，安装如下运行环境：

| 环境 | 版本 | 下载地址 |
|---|---|---|
| go | >= 1.19.2 | https://golang.google.cn/dl |
| mysql | >= 8.0.31 | https://www.mysql.com/downloads |
| redis | >= 7.0.5 | https://redis.io/download |
| node | >= 18.12.0 | https://nodejs.org/en/download |

## 运行服务端

::: tip 前提条件：
Go 环境正常、MySQL 服务正常启动、Redis 服务正常启动。
:::

### 导入数据库文件

您需要把 ```crm/server/db/crm.sql``` 文件导入到本地的数据库中，您可以在数据库开发工具 Navicat 中直接导入 SQL 文件。

或者使用 MySQL 自带的命令行工具导入 SQL 文件：

```bash
mysql> source 文件路径
```

### 修改配置文件

您需要修改 ```crm/server/config.yaml``` 配置文件，您可以保持默认配置，但是您必须修改文件存储配置、邮件服务配置和支付宝支付服务配置，否则您将无法正常使用文件导出、注册、密码找回、获取验证码、订阅、支付等功能。

文件存储配置如下：

```yaml
file:
  path: /home/ubuntu/crm/source/
```

把路径改成自己的系统路径，建议使用绝对路径。

邮件服务配置如下：

```yaml
# 邮件服务
mail:
  smtp: smtp.qq.com
  secret: dhsepilzlvoaceij   // 改成自己的QQ邮箱SMTP服务的授权码
  sender: 1655064994@qq.com  // 改成自己的QQ邮箱
```

**如何开启邮箱的 SMTP 服务并获取授权码？**

请参考官方文档[service.mail.qq.com](https://service.mail.qq.com)。

支付宝支付服务配置如下：

```yaml
# 支付宝支付服务配置
alipay:
  appId: 2022003122606990
  privateKey: MIIEpQIBAAKCAQEAkR0YofR...2sDd6uIy9rkpk8azj/rLmetW5r+tqTZgxcPWKeSz4=
  appPublicCert: /home/ubuntu/crm/cert/appPublicCert.crt
  alipayRootCert: /home/ubuntu/crm/cert/alipayRootCert.crt
  alipayPublicCert: /home/ubuntu/crm/cert/alipayPublicCert.crt
  returnURL: http://127.0.0.1:8000/api/subscribe/callback
  notifyURL: http://127.0.0.1:8000/api/subscribe/notify
  paySuccessURL: http://127.0.0.1:8060/#/subscribe
```

**如何获取支付宝支付服务的 appId、privateKey、appPublicCert、alipayRootCert、alipayPublicCert ？**

您需要登录到支付宝开放平台，请访问[open.alipay.com](https://open.alipay.com)。

登录到支付宝开放平台后，点击 ```控制台``` -> ```开发工具推荐``` -> ```沙箱``` -> ```沙箱应用```，在基本信息中可看到 ```APPID``` ，然后在开发信息中选择 ```系统默认秘钥``` -> ````启用证书模式``` -> ```查看```，即可看到复制私钥和下载证书，1个私钥和3个证书。

**如何获取支付账号？**

点击 ```控制台``` -> ```开发工具推荐``` -> ```沙箱``` -> ```沙箱账号```，里面有商家信息和买家信息，买家信息中的买家账号就是支付账号。

### 初始化并运行

使用电脑自带的终端执行执行如下命令，也可以使用 VS Code 或者 Goland 等开发工具，打开 crm 目录，找到 Terminal 终端。

执行如下命令：

```bash
$ cd server
$ go mod tidy
$ go build -o server main.go (windows编译命令为 go build -o server.exe main.go )

# 运行二进制
$ ./server (windows运行命令为 server.exe)
```

运行二进制文件后，如果控制台没有出现报错信息，就说明服务端启动成功。

## 运行Web端

::: tip 前提条件：
Node 环境正常。
:::

使用电脑自带的终端执行执行如下命令，也可以使用 VSCode 或者 WebStom 等开发工具，打开 crm 目录，找到 Terminal 终端。

执行如下命令：

```bash
$ cd web
$ npm install
$ npm run dev
```

当您使用 ```npm install``` 命令安装依赖时，如果出现下载缓慢、卡住等问题，您可能需要修改 npm 默认的镜像源。

要修改镜像源，请执行如下命令：

```bash
$ npm install -g cnpm --registry=https://registry.npmmirror.com
```

想要了解有关 Npm Mirror 中国镜像站的更多信息，请访问[npmmirror.com](https://npmmirror.com)。

当 Web 端启动成功后，打开浏览器访问http://127.0.0.1:8060。

## 部署到云服务器

以上是把项目部署到本地，如果您想部署到云服务器，请参考[部署到云服务器](/project/devops/deploy-cloudserver)。

