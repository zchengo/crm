# 部署到云服务器

如果您打算将 Crm 系统部署到您自己的云服务器，您需要先注册一个域名和购买一台云服务器，然后在云服务器上进行环境安装、服务启动与配置、项目的构建与部署。

## 注册域名

您可以选择以下任意云厂商提供的域名注册，完成域名注册。

- [阿里云-域名注册](https://wanwang.aliyun.com)
- [腾讯云-域名注册](https://dnspod.cloud.tencent.com/)
- [百度智能云-域名服务](https://cloud.baidu.com/product/bcd.html)

## 购买云服务器

您可以选择以下任意云厂商提供的云服务器，完成云服务器购买。

- [阿里云-云服务器ECS](https://www.aliyun.com/product/ecs)
- [腾讯云-云服务器CVM](https://cloud.tencent.com/product/cvm)
- [华为云-弹性云服务器ECS](https://www.huaweicloud.com/product/ecs.html)
- [亚马逊云-Amazon EC2](https://aws.amazon.com/cn/ec2/?nc2=h_ql_prod_fs_ec2)
- [百度智能云-云服务器BCC](https://cloud.baidu.com/product/bcc.html)

::: tip 友情提示
亚马逊云新用户可以免费体验12个月云服务器以及其他基础设施产品和服务，[前往免费体验](https://aws.amazon.com/cn/campaigns/freecenter-select-region)。
:::

## 创建云服务器实例

以亚马逊云 Amazon EC2（其他云服务器类似）为例，创建与配置云服务器实例。

第一步，登录到亚马逊云控制台，点击 ```EC2```。

<img src="/amazon_console.png" style="border-radius: 5px;" alt="amazon console"/>

第二步，选择 ```实例```，点击 ```启动新实例```。

<img src="/amazon_ec2.png" style="border-radius: 5px;" alt="amazon ec2"/>

第三步，填写实例名称 ```crm```（可填写任意名称），操作系统选择 ```Ubuntu```（比较推荐），其他按默认设置即可，然后点击 ```启动实例```。

<img src="/start_instance.png" style="border-radius: 5px;" alt="start instance"/>

第四步，实例启动成功后，您需要连接到您的实例，请点击 ```连接到实例```。

<img src="/start_instance_success.png" style="border-radius: 5px;" alt="start instance success"/>

第五步，连接到实例的方式有四种，推荐使用 ```SSH 客户端``` 来连接实例，您需要先关联密钥对，请根据如下提示进行操作。

<img src="/ssh_connect_ec2.png" style="border-radius: 5px;" alt="ssh connect ec2"/>

::: details 如何将密钥对与 EC2 实例（云服务器）进行关联？
1. 在您自己的电脑中，执行如下命令来生成密钥对：

```bash
ssh-keygen -t rsa -C "xxxx@qq.com"
```

生成密钥对后，会在 ```~/.ssh/``` 目录下生成私钥文件 ```id_rsa``` 和公钥文件 ```id_rsa.pub```，复制公钥文件中的内容。

2. 查看实例列表，选中您创建的实例，点击 ```连接```。

<img src="/instance_list.png" style="border-radius: 5px;" alt="instance list"/>

3. 在连接到实例页面中，选择 ```EC2 Instance Connect```，点击 ```连接```，即可连接到服务器。

4. 将自己电脑中生成的公钥文件的内容，复制到云服务器的 ```~/.ssh/authorized_keys``` 文件中。

完成以上步骤，即可完成密钥对与 EC2 实例（云服务器）的关联。
:::

密钥对与 EC2 实例（云服务器）关联成功后，您就可以使用 SSH 连接到您创建的 EC2 实例），您可以执行如下命令：

```bash
ssh 用户名@ip地址   // 比如 ssh ubuntu@52.223.50.216
```

进入到云服务器后，即可搭建系统的运行环境。

## 环境搭建

您需要安装以下运行环境：

| 环境 | 版本 | 下载地址 |
|---|---|---|
| go | >= 1.19.2 | https://golang.google.cn/dl |
| mysql | >= 8.0.31 | https://www.mysql.com/downloads |
| redis | >= 7.0.5 | https://redis.io/download |
| node | >= 18.12.0 | https://nodejs.org/en/download |
| nginx | >= 1.22.1 | http://nginx.org/en/download.html |

::: tip 提示
以下安装操作适用于 Ubuntu 操作系统，安装版本均为 Linux 版本。
:::

### 安装 Go 环境

安装 Go 语言压缩包，请执行如下命令：

```bash
$ wget https://golang.google.cn/dl/go1.19.5.linux-amd64.tar.gz
```

压缩包下载完成后，建议解压到 ```/usr/local/``` 目录下，您可以执行如下命令：

```bash
$ sudo tar -zxvf go1.19.5.linux-amd64.tar.gz -C /usr/local/
```

配置 Go 环境变量，使用 ```vi``` 命令打开 ```~/.profile``` 文件，添加如下配置：

```bash
# golang
export PATH="$PATH:/usr/local/go/bin"
```

添加完成后，执行 ```source .profile```。

查看 Go 环境是否安装成功：

```bash
$ go version
go version go1.19.5 linux/amd64
```

### 安装 Node 环境

下载 Node For Linux Binaries (x64) 并解压到 ```/usr/local/``` 目录下，请执行如下命令：

```bash
$ wget https://nodejs.org/dist/v18.13.0/node-v18.13.0-linux-x64.tar.xz
$ sudo tar -xvf node-v18.13.0-linux-x64.tar.xz -C /usr/local/
```

配置 Node 环境变量，使用 ```vi``` 命令打开 ```~/.profile``` 文件，添加如下配置：

```bash
# nodejs
export PATH="$PATH:/usr/local/node-v18.13.0-linux-x64/bin"
```

添加完成后，执行 ```source .profile```。

查看 Node 环境是否安装成功：

```bash
$ node -v
v18.13.0
$ npm -v
8.19.3
```

### 安装 MySQL

使用 ```apt install``` 安装，请执行如下命令：

```bash
$ sudo apt install mysql-server
```

查看是否安装成功：

```bash
$ mysql --version
mysql  Ver 8.0.31-0ubuntu0.22.04.1 for Linux on x86_64 ((Ubuntu))
```

默认用户名 ```root```, 默认密码 ``` ```，修改默认密码为 ```123456```，请执行如下命令：

```bash
# mysqladmin -u用户名 -p旧密码 password 新密码
$ sudo mysqladmin -uroot -p password 123456
```

登录到 MySQl，请执行如下命令：

```bash
$ sudo mysql -uroot -p123456
```

### 安装 Redis

安装 redis-stable 源码，请执行如下命令：

```bash
wget https://download.redis.io/redis-stable.tar.gz
```

编译和安装 redis，默认会安装到 ```/usr/local/bin``` 目录下，请执行如下命令：

```bash
$ sudo tar -xzvf redis-stable.tar.gz
$ cd redis-stable
$ sudo make
$ sudo make install
```

启动 redis 服务，请执行如下命令：

```bash
$ redis-server
```

### 安装 Nginx

安装 Nginx 源码包，请执行如下命令：

```bash
$ wget http://nginx.org/download/nginx-1.22.1.tar.gz
```

解压 ```nginx-1.22.1.tar.gz``` 到当前目录，请执行如下命令：

```bash
$ sudo tar -zxvf nginx-1.22.1.tar.gz
```

编译和安装 Nginx，请执行如下命令：

```bash
$ cd nginx-1.22.1
$ sudo ./configure --prefix=/usr/local/nginx --without-http_rewrite_module
```

::: danger 您可能会遇到如下报错：
```bash
./configure: error: the HTTP gzip module requires the zlib library.
You can either disable the module by using --without-http_gzip_module
option, or install the zlib library into the system, or build the zlib library
statically from the source with nginx by using --with-zlib=<path> option.
```
:::

需要先安装 ```zlib library```, 请执行如下命令：

```bash
$ wget http://www.zlib.net/zlib-1.2.13.tar.gz
$ tar -zxvf zlib-1.2.13.tar.gz
$ cd zlib-1.2.13
$ sudo ./configure
$ sudo make
$ sudo make install
```

了解 zlib 的更多信息，请访问[zlib.net](http://www.zlib.net/)。

安装完 ```zlib library``` 后，重新安装 Nginx，请执行如下命令：

```bash
$ cd nginx-1.22.1
$ sudo ./configure --prefix=/usr/local/nginx --without-http_rewrite_module
$ sudo make
$ sudo make install
```

查看是否安装成功，请执行如下命令：

```bash
$ nginx -v
nginx version: nginx/1.22.1
```

修改 Nginx 配置文件 ```/usr/local/nginx/conf/nginx.conf```，请参考如下配置：

```bash
#user  nobody;
worker_processes  1;

#error_log  logs/error.log;
#error_log  logs/error.log  notice;
#error_log  logs/error.log  info;

#pid        logs/nginx.pid;

events {
    worker_connections  1024;
}

http {
    include       mime.types;
    default_type  application/octet-stream;

    #log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
    #                  '$status $body_bytes_sent "$http_referer" '
    #                  '"$http_user_agent" "$http_x_forwarded_for"';

    #access_log  logs/access.log  main;

    sendfile        on;
    #tcp_nopush     on;

    #keepalive_timeout  0;
    keepalive_timeout  65;

    #gzip  on;

    server {
        listen       80;
        server_name  zocrm.cloud;

        #charset koi8-r;

        #access_log  logs/host.access.log  main;

        location / {
            root   html/dist;
            index  index.html;
        }

        #error_page  404              /404.html;

        # redirect server error pages to the static page /50x.html
        #
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }

	    # api proxy
	    location /api {
            proxy_pass http://127.0.0.1:8000;
        }

        # proxy the PHP scripts to Apache listening on 127.0.0.1:80
        #
        #location ~ \.php$ {
        #    proxy_pass   http://127.0.0.1;
        #}

        # pass the PHP scripts to FastCGI server listening on 127.0.0.1:9000
        #
        #location ~ \.php$ {
        #    root           html;
        #    fastcgi_pass   127.0.0.1:9000;
        #    fastcgi_index  index.php;
        #    fastcgi_param  SCRIPT_FILENAME  /scripts$fastcgi_script_name;
        #    include        fastcgi_params;
        #}

        # deny access to .htaccess files, if Apache's document root
        # concurs with nginx's one
        #
        #location ~ /\.ht {
        #    deny  all;
        #}
    }


    # another virtual host using mix of IP-, name-, and port-based configuration
    #
    #server {
    #    listen       8000;
    #    listen       somename:8080;
    #    server_name  somename  alias  another.alias;

    #    location / {
    #        root   html;
    #        index  index.html index.htm;
    #    }
    #}


    # HTTPS server
    
    server {
        listen       443 ssl;
        server_name  zocrm.cloud;

        ssl_certificate      cert/8913125_zocrm.cloud.pem;
        ssl_certificate_key  cert/8913125_zocrm.cloud.key;

    #    ssl_session_cache    shared:SSL:1m;
    #    ssl_session_timeout  5m;

    #    ssl_ciphers  HIGH:!aNULL:!MD5;
    #    ssl_prefer_server_ciphers  on;

        location / {
            root   html/dist;
            index  index.html;
        }
	    location /api {
            proxy_pass http://127.0.0.1:8000;
        }
    }

    server {
        listen       443 ssl;
        server_name  docs.zocrm.cloud;

        ssl_certificate      cert/9129982_docs.zocrm.cloud.pem;
        ssl_certificate_key  cert/9129982_docs.zocrm.cloud.key;

    #    ssl_session_cache    shared:SSL:1m;
    #    ssl_session_timeout  5m;

    #    ssl_ciphers  HIGH:!aNULL:!MD5;
    #    ssl_prefer_server_ciphers  on;

        location / {
            root   html/docs/dist;
            index  index.html;
        }
    }
}
```

启动 Nginx 服务，请执行如下命令：

```bash
$ cd /usr/local/nginx/sbin/
$ sudo ./nginx -c /usr/local/nginx/conf/nginx.conf
# 或者
$ sudo nginx
```

当 Nginx 服务成功启动后，您可以尝试在浏览器中访问服务器的 ```IP地址```，查看是否能访问到 Nginx 的默认页面。

::: details 访问不到 Nginx 默认页面？
您可能需要配置服务器实例的安全组，进入到 ```Amazon EC2 控制台``` -> ```EC2``` -> ```安全组``` -> ```入站规则``` -> ```编辑入站规则信息``` -> ```添加规则```，请参考如下：

<img src="/incoming_rules.png" style="border-radius: 5px;" alt="incoming rules"/>

点击 ```保存规则```，在浏览器中访问 ```IP地址```，即可看到如下页面：

<img src="/nginx_home.png" style="border-radius: 5px;" alt="nginx home page"/>
:::

到这里，系统运行环境的搭建基本已经完成。

## 构建与部署

您需要先下载 Crm 项目代码到服务器。

如果您的系统中已经安装了 ```Git```，您可以使用如下方式安装（推荐）：

```bash
git clone https://github.com/zchengo/crm.git
```

您也可以使用 ```wget``` 命令下载压缩包，使用 ```unzip``` 解压，请执行如下命令：

```bash
wget https://github.com/zchengo/crm/archive/refs/heads/main.zip
unzip main.zip
```

### 部署服务端

::: tip 前提条件：
Go 环境正常、MySQL 服务正常启动、Redis 服务正常启动。
:::

#### 导入数据库文件

请先登录到 MySQl，请执行如下命令：

```bash
$ sudo mysql -uroot -p123456
```

再把 ```crm/server/db/crm.sql``` 文件导入到 MySQL 数据库中，执行如下 SQL 语句：

```bash
mysql> source SQL文件路径
```

#### 修改配置文件

您需要修改 ```crm/server/config.yaml``` 配置文件，必须修改文件存储配置、邮件服务配置和支付宝支付服务配置，否则您将无法正常使用文件导出、注册、密码找回、获取验证码、订阅、支付等功能。

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

将 ```returnURL```、```notifyURL```、```paySuccessURL``` 中的 ```IP地址```改成您的服务器的 ```IP地址``` 或 ```域名```。

#### 初始化和运行

服务端初始化和运行，请执行如下命令：

```bash
$ cd crm
$ cd server
$ go mod tidy
$ go build -o crmserver main.go (windows编译命令为 go build -o crmserver.exe main.go )

# 运行二进制
$ ./crmserver (windows运行命令为 crmserver.exe)
# 或者在后台运行
$ sudo nohup ./crmserver > /home/ubuntu/crmserver.log 2>&1 &
```

### 部署Web端

::: tip 前提条件：
Node 环境正常。
:::

Web端初始化和打包，请执行如下命令：

```bash
$ cd crm
$ cd web
$ npm install
$ npm run build
```

::: details 使用 ```npm install``` 命令安装依赖时，出现下载缓慢、卡住等问题?
您可能需要修改 npm 默认的镜像源，请执行如下命令：

```bash
$ npm install -g cnpm --registry=https://registry.npmmirror.com
```

想要了解有关 Npm Mirror 中国镜像站的更多信息，请访问[npmmirror.com](https://npmmirror.com)。
:::

打包完成后，会在 ```crm/web/dist/``` 目录下生成页面资源文件，您需要复制这些文件到 ```/usr/local/nginx/html/``` 目录下，请执行如下命令：

```bash
sudo cp ./crm/web/dist /usr/local/nginx/html/
```

您可能需要适当的修改 Nginx 配置文件 ```/usr/local/nginx/conf/nginx.conf```，以确保能够访问到页面，您可能需要修改这个部分：

```bash
http {
    ...
    server {
        ...
        location / {
            root   html/dist;
            index  index.html;
        }
    }
}
```

然后您就可以通过 ```IP地址``` 或 ```域名``` 来访问 Crm 系统的页面了。

### 部署Docs文档

如果您有需要，您也可以部署 Docs 项目文档，部署方式与部署Web端类似。

请参考如下命令：

```bash
$ cd crm/docs
$ npm install
$ npm run docs:build
$ sudo cp ./docs/docs/.vitepress/dist /usr/local/nginx/html/docs
```

修改 Nginx 配置文件请参考：

```bash
#user  nobody;
worker_processes  1;

#error_log  logs/error.log;
#error_log  logs/error.log  notice;
#error_log  logs/error.log  info;

#pid        logs/nginx.pid;

events {
    worker_connections  1024;
}

http {
    
    ...

    # HTTPS server
    
    # Web
    server {
        listen       443 ssl;
        server_name  zocrm.cloud;

        ssl_certificate      cert/8913125_zocrm.cloud.pem;
        ssl_certificate_key  cert/8913125_zocrm.cloud.key;

        location / {
            root   html/dist;
            index  index.html;
        }
	    location /api {
            proxy_pass http://127.0.0.1:8000;
        }
    }

    # Docs
    server {
        listen       443 ssl;
        server_name  docs.zocrm.cloud;

        ssl_certificate      cert/9129982_docs.zocrm.cloud.pem;
        ssl_certificate_key  cert/9129982_docs.zocrm.cloud.key;

        location / {
            root   html/docs/dist;
            index  index.html;
        }
    }
}
```

当然，如果你使用域名来替换 ```html/``` 目录下的 ```dist``` 目录名称，或许会更优雅。

比如您可以这样命名：

```text
/usr/local/nginx/html/www/zocrm/com
/usr/local/nginx/html/www/docs/zocrm/com
```

### 测试

在浏览器访问您部署的项目，比如 http://3.83.115.236 或者 http://zocrm.cloud。

如果您配置了 SSL 证书，您可以使用 HTTPS 来访问。

## 持续集成与交付

当有新版本发布时，即可自动部署到服务器，请参考[持续集成与交付 (CI/CD)](/project/devops/ci-cd)。

## 常见问题

如果您在部署过程中遇到了问题，您可以通过以下方式反馈：

- [在公众号后台反馈](/about/about)
- [New Issues In Github](https://github.com/zchengo/crm/issues)

## 支持作者

此文档对您有帮助的话，您可以[支持作者](/sponsor/sponsor)。















